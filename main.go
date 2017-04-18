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
	if len(*sqlconn) == 0 {
		log.Fatal("connection string not provided")
	}

	r, err := db.New(*sqlconn)
	if err != nil {
		log.Fatal(err)
	}

	var a *api.Api
	a, err = api.New(r)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(a.Serve())
}
