package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mux_test/utils"
	"net/http"
)

func HostApp() {
	utils.ConnectDB()
	router := mux.NewRouter()
	Map(router)
	fmt.Println("starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
