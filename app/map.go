package app

import (
	"github.com/gorilla/mux"
	"mux_test/handler"
)

func Map(router *mux.Router) {
	router.HandleFunc("/signup", handler.SignupHandler).Methods("POST")
	router.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	router.HandleFunc("/protected", handler.TokenVerifyingMiddleware(handler.ProtectedHandler)).Methods("GET")
}
