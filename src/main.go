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

func (handler HttpHandler) ListenAndServe(port string) {
	err := http.ListenAndServe(":8080", handler)
	
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}

func main() {
	kernal := HttpHandler{}
	kernal.ListenAndServe(":8080")
}