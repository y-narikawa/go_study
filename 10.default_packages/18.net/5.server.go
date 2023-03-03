package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "Hello")
}

func main() {

	http.ListenAndServe(":8080", &MyHandler{})

	// not found

	// todo 壊れてるからいつかなおす
	http.HandleFunc("/top", top)
	http.ListenAndServe(":8080", nil)
}