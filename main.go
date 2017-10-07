package main

import (
	"os"
	"log"
	"net/http"
	"fmt"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	port:= os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port not set")
	}

	server := http.Server{
		Addr: ":" + port,
	}

	http.HandleFunc("/", index)
	server.ListenAndServe()
}