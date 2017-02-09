package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("ok")
}

func serve() {
	// TODO: Pull listen interface and port from a flag
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func registerRoutes() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Printf("%s\t%s\n", req.Method, req.URL)
		switch req.Method {
		case "GET":
			GETguestBook(res, req)
		case "POST":
			POSTguestBook(res, req)
		}
	})
}
