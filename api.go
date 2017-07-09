package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type someThing struct {
	fName string
	fAge  int
}

func (p *someThing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, p.fName)
	log.Println(r.Method)
	SomeFunc(w, r)

}
func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	//someOne := &someThing{"test123213", 30}
	http.HandleFunc("/", notFound)
	http.HandleFunc("/api", testAPI)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Cai gi cung dc")
	log.Println(r.Method + " - " + r.URL.Path)
}
func testAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Method)
}
