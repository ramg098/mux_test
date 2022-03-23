package handler

import "net/http"

func SignupHandler(v http.ResponseWriter, r *http.Request) {
	v.Write([]byte("Signup invoked"))
}

func LoginHandler(v http.ResponseWriter, r *http.Request) {
	v.Write([]byte("Login invoked"))
}

func ProtectedHandler(v http.ResponseWriter, r *http.Request) {
	v.Write([]byte("Protected invoked"))
}

func TokenVerifyingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	//ProtectedHandler()
	return next
}
