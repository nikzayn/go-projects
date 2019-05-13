package main

import (
	"fmt"
	"net/http"
)

type some int

func (m some) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Nik", "This side is Batman")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "<h1>Anything you want</h1>")
}

func main() {
	var d some
	http.ListenAndServe(":8080", d)
}
