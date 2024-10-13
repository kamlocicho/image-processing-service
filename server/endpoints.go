package server

import (
	"fmt"
	"log"
	"net/http"
)

func setupEndpoints() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server running.")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
