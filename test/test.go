package test

import (
	"fmt"
	"log"
	"net/http"
)

type someThing struct {
	fName string
	fAge  int
}

func (p *someThing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, p.fName)
	log.Println(r.Method)

}
func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not Found 404")
	log.Println(r.Method + " - " + r.URL.Path)
}

// SomeFunc formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
func SomeFunc(w http.ResponseWriter, r *http.Request) {
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
