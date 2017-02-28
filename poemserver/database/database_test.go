package database_test

import (
	. "poemXML/poemserver/database"
	"testing"
)

func TestConstructDBConnectionString(t *testing.T) {
	user := "root"
	password := "potato"
	host := "127.0.0.1"
	dbName := "hello"
	dbString := ConstructDBConnectionString(host, user, dbName, password)
	expected := "root:potato@tcp(127.0.0.1:3306)/hello"

	if dbString != expected {
		t.Errorf("Expected %s, got %s", expected, dbString)
	}
}
