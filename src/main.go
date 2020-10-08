package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// type ServeMux struct {}

type HttpHandler struct {}

func (handler HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// data := []byte("Hello World")

	// write `Hello` using `io.WriteString` function
    io.WriteString(res, "Hello")
    // write `World` using `fmt.Fprint` function
    fmt.Fprint(res, " World! ")
    // write `❤️` using simple `Write` call
    res.Write([]byte("❤️"))
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, `
			<h1>Go HTTP w/ Docker</h1>
			<a href="/hello">Hello</a>
			<a href="/error">Error</a>
		`)
	})
	http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		io.WriteString(res, "Hello World!")
	})
	http.HandleFunc("/error", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusBadRequest)
		io.WriteString(res, `{"status": "ERROR"}`)
	})

	err := http.ListenAndServe(":8080", nil)
	
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}