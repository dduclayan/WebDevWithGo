package main

import (
	"fmt"
	"net/http"
)

func pathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
}

func main() {
	http.HandleFunc("/", pathHandler) // "localhost:3000/dog/cat/pig" will print "/dog/cat/pig"
	//http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", nil)
}
