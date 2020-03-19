package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	name := "World"

	param, ok := r.URL.Query()["name"]
	if ok {
		name = param[0]
	}

	io.WriteString(w, "Hello "+name+"!")
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8080", nil)
}
