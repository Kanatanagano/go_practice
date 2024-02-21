package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// analyze queryparameters
	query := r.URL.Query()
	name := query.Get("name")

	// create a response map
	response := map[string]string{
		"message": "Hello " + name,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	fmt.Println("Starting server on port 8010")
	http.HandleFunc("/api/hello", helloHandler)
	http.ListenAndServe(":8010", nil)
}
