package main

import (
	"flag"
	"log"

	"github.com/ContainerSolutions/guestbook/api"
	"github.com/ContainerSolutions/guestbook/db"
)

func main() {
	sqlconn := flag.String("postgres", "", "connect string for postgres")
	flag.Parse()
	r, err := db.New(*sqlconn)
	if err != nil {
		log.Fatal(err)
	}
	//r := entries.Retriever(entries.Mocker{})
	var a *api.Api
	a, err = api.New(r)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(a.Serve())

}
