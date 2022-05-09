package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email at <a href=\"mailto:devin@duclayan.com\">"+
		"devin@duclayan.com</a>.")
}

func faqHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, `<h1>FAQ</h1>
	<ul>
		<li><b>Where are you from?</b> Hawaii!</li>
		<li><b>What is your favorite vacation destination?</b> Japan!</li>
		<li><b>What is your favorite hobby?</b> Mechanical Keyboards :) </li>	
	<ul>
`)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(http.NotFound)
	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", r)
}
