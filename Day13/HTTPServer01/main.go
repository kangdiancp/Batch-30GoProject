package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("My first web with go"))
}

func hello(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello World\n")
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/hello", hello)

	log.Println("Starting HTTP Server")
	//start an HTTP Server at port 8888
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server already running at port :8888")
}
