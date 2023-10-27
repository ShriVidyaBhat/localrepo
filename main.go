package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//connect
	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks)

	log.Fatal(http.ListenAndServe(":90", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {

}
