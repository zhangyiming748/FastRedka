package db

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/nalgeon/redka"
)

var (
	db *redka.DB
)

func GetDatabase() *redka.DB {
	return db
}
func Initial() {
	db, _ = redka.Open("data.db", nil)
}
