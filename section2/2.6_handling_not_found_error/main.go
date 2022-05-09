package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email at <a href=\"mailto:devin@duclayan.com\">"+
		"devin@duclayan.com</a>.")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.NotFound(w, r)
		// alternative methods below
		//http.Error(w, http.StatusText(404), 404)
		//http.Error(w, "Page not found", http.StatusNotFound)
		//w.WriteHeader(http.StatusNotFound)
		//fmt.Fprint(w, "Page not found")
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", nil)
}
