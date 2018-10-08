package model

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func ConnectDb(user string, password string, database string, address string) (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
		Addr:     address,
	})

	return db
}

func MigrationDb(db *pg.DB, schema string) error {
	// Tạo schema theo tên service
	_, err := db.Exec("CREATE SCHEMA IF NOT EXISTS " + schema + ";")
	if err != nil {
		return err
	}

	// Tạo bảng
	var user User
	err = createTable(&user, db)
	if err != nil {
		return err
	}

	return nil
}

func LogQueryToConsole(db *pg.DB) {
	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}

		log.Printf("%s %s", time.Since(event.StartTime), query)
	})
}

func createTable(model interface{}, db *pg.DB) error {
	err := db.CreateTable(model, &orm.CreateTableOptions{
		Temp:          false,
		FKConstraints: true,
		IfNotExists:   true,
	})

	return err
}
