/*
Ex1 - Add a URL Parameter

Read the docs and see if you can add a URL parameter to one of your routes, retrieve it in your handler, and output it to the resulting HTML.

For example, imagine that we wanted to add support for a path used to show individual galleries.
We might use the path /galleries/<id-here>.
See if you can create a handle that retrieves the id from the URL using Chi, and set up the Chi router to direct requests to the correct handler.

	Hint:
	See these docs[https://github.com/go-chi/chi#url-parameters] if you need some guidance.
	You shouldn’t need to use context, just the URLParam method.

Ex2 - Experiment with Chi’s builtin middleware

Chi provides quite a few builtin middleware. One is the Logger middleware, which will track how long each request is taking.
Try to add it to your application, then to only a single route.

What other types of uses could you imagine middleware serving in an application?
*/

package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
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

func paramHandler(w http.ResponseWriter, r *http.Request) {
	v := chi.URLParam(r, "name")
	if v == "" {
		fmt.Fprintf(w, "<p>hello, NoName!</p>")
	} else {
		fmt.Fprintf(w, "<p>hi %v!</p>", v)
	}
}

func main() {
	r := chi.NewRouter()
	// r.Use(middleware.Logger)  // for entire app

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Route("/say", func(r chi.Router) {
		r.Use(middleware.Logger) // for single route
		r.Get("/{name}", paramHandler)
		r.Get("/", paramHandler)
	})
	r.NotFound(http.NotFound)

	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", r)
}
