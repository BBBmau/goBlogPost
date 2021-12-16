package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tplFolder embed.FS

//var blogPage []byte

func indexHanlder(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(tplFolder, "templates/blogPage.html")
	if err != nil {
		log.Fatal(err)
	}
	println("Embedded the blogPage File")
	tmpl.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHanlder)

	fmt.Printf("Starting server at port 8000\n")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
