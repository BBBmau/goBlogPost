package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/homePage.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
	fmt.Printf("In homeHandler!\n")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/aboutPage.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
	fmt.Printf("In aboutHandler!\n")
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/blogPage.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
	fmt.Printf("In blogHandler!")
}

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/home", homeHandler)
	route.HandleFunc("/about", aboutHandler)
	route.HandleFunc("/blog", blogHandler)
	fmt.Printf("Server up!\n")

	err := http.ListenAndServe(":8000", route)

	if err != nil {
		log.Fatal(err)
	}

}
