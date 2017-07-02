package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type someThing struct {
	fName string
}

func (p *someThing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, p.fName)
	log.Println(r.Method)
}
func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	someOne := &someThing{"test123213"}

	http.HandleFunc("/", notFound)
	http.HandleFunc("/api", someOne.ServeHTTP)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not Found 404")
	log.Println(r.Method + " - " + r.URL.Path)
}
func testAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintln(w, r.Method)
	}
	if r.Method == "POST" {
		fmt.Fprintln(w, r.Method)
	}
	if r.Method == "PUT" {
		fmt.Fprintln(w, r.Method)
	}
}
