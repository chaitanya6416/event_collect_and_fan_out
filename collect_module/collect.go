package main

import (
	"fmt"
	"io"
	"net/http"
)

func handleData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		fmt.Printf("Received data: %s\n", body)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handleData)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
