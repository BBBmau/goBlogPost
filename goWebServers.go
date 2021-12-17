package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/blogPage.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
	fmt.Printf("In Handler!")
}

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/blogPage", handler)
	fmt.Printf("Server up!\n")

	err := http.ListenAndServe(":8000", route)

	if err != nil {
		log.Fatal(err)
	}

}
