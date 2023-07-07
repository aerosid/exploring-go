package db

import (
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

var lock *sync.Mutex = &sync.Mutex{}

func GetInstance() (*sqlx.DB, error) {
	lock.Lock()
	defer lock.Unlock()
	var err error = nil
	if db == nil {
		db, err = sqlx.Connect("mysql", getDSN())
	}
	return db, err
}

func getDSN() (dsn string) {
	cfg := mysql.Config{
		User:   "go",
		Passwd: "hello",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "sample",
	}
	dsn = cfg.FormatDSN()
	return dsn
}
