package main // import "github.com/xmlking/go_project"

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
}

func loveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/love", loveHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server starting on port: 8080")
	http.ListenAndServe(":8080", nil)
}
