package db

import (
	"log"
	"log/slog"
	"testing"
)

func init() {
	Initial()
}
func TestExam(t *testing.T) {
	d := GetDatabase()
	// Open a database.

	// Set some string keys.
	key1 := "name"
	value1 := "alice"
	err := d.Str().Set(key1, value1)
	log.Printf("set key = %v\tvalue = %v\twith error:%v\n", key1, value1, err)
	key2 := "age"
	value2 := 25
	err = d.Str().Set(key2, value2)
	log.Printf("set key = %v\tvalue = %v\twith error:%v\n", key2, value2, err)

	// Check if the keys exist.
	count, err := d.Key().Count("name", "age", "city")
	slog.Info("count", "count", count, "err", err)

	// Get a key.
	name, err := d.Str().Get("name")
	slog.Info("get", "name", name, "err", err)
}
