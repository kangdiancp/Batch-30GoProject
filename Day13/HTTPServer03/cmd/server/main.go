package main

import (
	router "day13/HTTPServer03/cmd/routers"
	"day13/HTTPServer03/internal/store"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	router := router.NewRouter()
	store.DBConnect()

	//start an HTTP Server at port 8888
	err := http.ListenAndServe(":8888", router)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Server already running")

	//log.Fatal(http.ListenAndServe(":8888", router))

}
