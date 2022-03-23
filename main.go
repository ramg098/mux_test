package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/signup", SignupHandler).Methods("POST")
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyingMiddleware(ProtectedHandler)).Methods("GET")

	fmt.Println("starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func SignupHandler(v http.ResponseWriter, r *http.Request) {
	v.Write([]byte("Signup invoked"))
}

func LoginHandler(v http.ResponseWriter, r *http.Request) {
	v.Write([]byte("Signup invoked"))
}

func ProtectedHandler(v http.ResponseWriter, r *http.Request) {
	v.Write([]byte("Signup invoked"))
}

func TokenVerifyingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	//ProtectedHandler()
	return nil
}
