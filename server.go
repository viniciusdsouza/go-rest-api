package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Up and running...")
	})
	router.HandleFunc("/posts", GetPosts).Methods("GET")
	router.HandleFunc("/posts", AddPosts).Methods("POST")
	log.Println("Server lintening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
