package db

import (
	"database/sql"

	"github.com/ContainerSolutions/guestbook/entries"
	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

const (
	tablename = "entries"
)

type DB struct {
	dbMap *gorp.DbMap
}

func New(connstr string) (entries.Retriever, error) {
	db := &DB{}
	err := db.init(connstr)
	return db, err

}

func (d *DB) GetAll() ([]entries.Entry, error) {
	var es []entries.Entry
	_, err := d.dbMap.Select(&es, "select * from ? order by Date", tablename)
	return es, err
}

func (d *DB) Write(e entries.Entry) error {
	return d.dbMap.Insert(e)
}

func (d *DB) init(connstr string) error {
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return err
	}
	d.dbMap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	d.dbMap.AddTableWithName(entries.Entry{}, tablename)
	return d.dbMap.CreateTablesIfNotExists()
}
