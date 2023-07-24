package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", servceHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r)) // open port 4000 and route serveHome
}

func greeter() {
	fmt.Println("Hi Module port is open localhost:4000")
}

func servceHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to goland mod<h1>"))
}
