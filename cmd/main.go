package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	}

	r.HandleFunc("/", handler).Methods("GET")

	err := http.ListenAndServe(":80", r)
	if err != nil {
		fmt.Println(err)
	}
}