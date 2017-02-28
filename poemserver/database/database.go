package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
)

const PORT = "3306"

//DB Config

type DBConfig struct {
	host       string
	user       string
	dbName     string
	dbPassword string
}

func NewDBConfig(host, user, dbName, dbPassword string) *DBConfig {
	return &DBConfig{host: host, user: user, dbName: dbName, dbPassword: dbPassword}
}

//DB Connection

type Database struct {
	conn   mysql.Conn
	config *DBConfig
}

func ConstructDBConnectionString(host, dbUser, dbName, dbPassword string) string {
	return dbUser + ":" + dbPassword + "@tcp(" + host + ":" + PORT + ")/" + dbName

}

func NewConnectedDB(config *DBConfig) (db *sql.DB, err error) {
	dbString := ConstructDBConnectionString(config.host, config.user, config.dbName, config.dbPassword)

	db, err = sql.Open("mysql", dbString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}

func (db *Database) connect() (err error) {
	return nil
}
