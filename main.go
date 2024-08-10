package main

import (
	"fmt"
	"log"
	"net/http"

	ascii "ascii/Handlers"
)

func main() {
	http.HandleFunc("/", ascii.IndexHandler)
	http.HandleFunc("/ascii-art", ascii.AsciiHandler)
	// This creates a file server to serve static files from the stylize directory.
	http.HandleFunc("/css/", Style)

	fmt.Println("Server starting on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Style(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/css/" {
		ascii.ErrorHandler(w, "404 : Not Found", http.StatusNotFound)
		return
	}
	// this serve static files and it strips the /css prefix and serves the static file from that directory
	stylize := http.FileServer(http.Dir("stylize"))
	http.StripPrefix("/css", stylize).ServeHTTP(w, r)
}
