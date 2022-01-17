package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home page!")
}

func main() {

	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", homePage).Methods("GET")

	port := os.Getenv("PORT")
	// log.Fatal(http.ListenAndServe(":3000", myrouter)) //for testing on local machine
	log.Fatal(http.ListenAndServe(":"+port, myrouter))
}
