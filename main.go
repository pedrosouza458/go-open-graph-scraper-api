package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/websites/", websiteHandler)
	http.ListenAndServe(":8080", nil)
}

func websiteHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the URL path
	path := strings.TrimPrefix(r.URL.Path, "/websites/")
	if path == "" {
		http.Error(w, "User ID not provided", http.StatusBadRequest)
		return
	}

	// Respond with the user ID
	fmt.Fprintf(w, "User ID: %s", path)
}
