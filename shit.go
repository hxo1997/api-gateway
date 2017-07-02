package main

import (
	"fmt"
	"net/http"
)

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
