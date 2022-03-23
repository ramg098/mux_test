package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HostApp() {
	router := mux.NewRouter()
	Map(router)
	fmt.Println("starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
