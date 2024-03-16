package main

import (
	"net/http"
)

func main() {
	// Serve files from the "static" directory
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Start the server
	http.ListenAndServe(":8080", nil)
}
