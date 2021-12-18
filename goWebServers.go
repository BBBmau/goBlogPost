package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
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

//go:embed templates/*
var tp1Folder embed.FS

//go:embed assets/*
var assets embed.FS

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFS(tp1Folder, "templates/homePage.html")
	if err != nil {
		log.Fatal(err)
	}
	println("Embedded homePage.html")
	tmpl.Execute(w, nil)
}

func main() {
	assetsFS := http.FileServer(http.FS(assets))
	route := http.NewServeMux()
	route.Handle("/assets/", assetsFS)
	route.HandleFunc("/home", homeHandler)
	route.HandleFunc("/about", aboutHandler)
	route.HandleFunc("/blog", blogHandler)

	//route.Handle("comma.png")

	route.HandleFunc("/", indexHandler)
	fmt.Printf("Server up!\n")

	err := http.ListenAndServe(":8000", route)

	if err != nil {
		log.Fatal(err)
	}

}
