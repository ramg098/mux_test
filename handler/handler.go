package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"mux_test/model"
	"net/http"
	"strings"
)

func SignupHandler(v http.ResponseWriter, r *http.Request) {
	var user model.User
	var error model.Error
	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "email is missing"
		respondWithError(v, http.StatusBadRequest, error)
		return
	}
	if user.Password == "" {
		error.Message = "password is missing"
		respondWithError(v, http.StatusBadRequest, error)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(hash)
	stmt := "insert into users (email, password) values($1, $2) RETURNING id;"
	err = model.DB.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		error.Message = "Server error"
		respondWithError(v, http.StatusInternalServerError, error)
		return
	}

	user.Password = ""
	v.Header().Set("Content-Type", "application/json")
	responseJSON(v, user)
}

func LoginHandler(v http.ResponseWriter, r *http.Request) {
	var user model.User
	var jwt model.JWT
	var error model.Error
	json.NewDecoder(r.Body).Decode(&user)
	if user.Email == "" {
		error.Message = "email is missing"
		respondWithError(v, http.StatusBadRequest, error)
		return
	}
	if user.Password == "" {
		error.Message = "password is missing"
		respondWithError(v, http.StatusBadRequest, error)
		return
	}

	password := user.Password
	row := model.DB.QueryRow("select * from users where email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			error.Message = "User not exist"
			respondWithError(v, http.StatusBadRequest, error)
			return
		} else {
			log.Fatal("3", err)
		}
	}
	hashedPassword := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		error.Message = "Invalid Password"
		respondWithError(v, http.StatusUnauthorized, error)
		return
	}
	token, err := GenerateToken(user)
	if err != nil {
		log.Fatal(err)
	}
	v.WriteHeader(http.StatusOK)
	jwt.Token = token

	responseJSON(v, jwt)
}

func ProtectedHandler(v http.ResponseWriter, r *http.Request) {
	v.Write([]byte("success"))
}

func TokenVerifyingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(v http.ResponseWriter, r *http.Request) {
		var errObject model.Error
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			authToken := bearerToken[1]
			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf(" an error")
				}
				return []byte("secret"), nil
			})
			if error != nil {
				errObject.Message = error.Error()
				respondWithError(v, http.StatusUnauthorized, errObject)
				return
			}

			if token.Valid {
				next.ServeHTTP(v, r)
			} else {
				errObject.Message = error.Error()
				respondWithError(v, http.StatusUnauthorized, errObject)
				return
			}
		} else {
			errObject.Message = "Invalid Token"
			respondWithError(v, http.StatusUnauthorized, errObject)
			return
		}
	})
}

func respondWithError(v http.ResponseWriter, status int, error model.Error) {
	v.WriteHeader(status)
	json.NewEncoder(v).Encode(error)
}

func responseJSON(v http.ResponseWriter, data interface{}) {
	json.NewEncoder(v).Encode(data)
}

func GenerateToken(user model.User) (string, error) {
	var err error
	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString, err
}
