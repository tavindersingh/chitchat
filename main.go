package main

import (
	"os"
	"log"
	"net/http"
	"html/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hello World!")
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func main() {
	port:= os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port not set")
	}

	server := http.Server{
		Addr: ":" + "3001",
	}

	http.HandleFunc("/", index)
	server.ListenAndServe()
}