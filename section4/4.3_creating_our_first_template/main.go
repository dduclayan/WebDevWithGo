// You need to be in the dir where hello.gohtml lives when running 'go run main.go'

package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta struct {
		Visits int
	}
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Susan Smith",
		Age:  12,
		Meta: UserMeta{
			Visits: 4,
		},
	}

	// example of an anonymous struct
	//user := struct {
	//	Name string
	//	Age  int
	//}{
	//	Name: "Susan Smith",
	//	Age:  12,
	//}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
