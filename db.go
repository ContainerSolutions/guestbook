package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	createdbQuery    = `CREATE DATABASE guestbook;`
	createtableQuery = `CREATE TABLE messages (message text, created_at timestamp DEFAULT current_timestamp);`
)

func conn(connstr string) (*sql.DB, error) {
	return sql.Open("postgres", connstr)
}

func createdb(db *sql.DB) error {
	res, err := db.Exec(createdbQuery)
	if err != nil {
		return err
	}

	rowsAf, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Rows affected: %d\n", rowsAf)
	return nil
}

func createtable(db *sql.DB) error {
	res, err := db.Exec(createtableQuery)
	if err != nil {
		return err
	}

	rowsAf, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Rows affected: %d\n", rowsAf)
	return nil
}
