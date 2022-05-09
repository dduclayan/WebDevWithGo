/*
Add an FAQ page.

We want to visit <localhost:3000/faq> and have it render an FAQ page with questions and answers like:
```
Q: Where are you from?
A: Hawaii!

Q: What is favorite vacation destination?
A: Japan!

Q: What is your favorite hobby?
A: Nothing!
```
*/

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

func faqHandler(w http.ResponseWriter, _ *http.Request) {
	//fmt.Fprint(w, "<h1>FAQ</h1>")
	//fmt.Fprint(w, "<p>Q: Where are you from?</p>")
	//fmt.Fprint(w, "<p>A: Hawaii!</p>")
	//fmt.Fprint(w, "<br></br>")
	//fmt.Fprint(w, "<p>Q: What is your favorite vacation destination?</p>")
	//fmt.Fprint(w, "<p>A: Japan!</p>")
	//fmt.Fprint(w, "<br></br>")
	//fmt.Fprint(w, "<p>Q: What is your favorite hobby?</p>")
	//fmt.Fprint(w, "<p>A: Mechanical Keyboards :)</p>")
	//fmt.Fprint(w, "<br></br>")

	fmt.Fprint(w, `<h1>FAQ</h1>
<ul>
	<li><b>Where are you from?</b> Hawaii!</li>
	<li><b>What is your favorite vacation destination?</b> Japan!</li>
	<li><b>What is your favorite hobby?</b> Mechanical Keyboards :) </li>	
<ul>
`)
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	var router Router
	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", router)
}
