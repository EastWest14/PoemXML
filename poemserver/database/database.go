package database

import (
	"fmt"
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

func NewConnectedDB(config *DBConfig) (db *Database, err error) {
	db = &Database{config: config}
	fmt.Println("Connecting to DB")
	err = db.connect()
	if err != nil {
		return nil, err
	}

	connection := mysql.New("tcp", "", config.host+":"+PORT, config.user, config.dbPassword, config.dbName)
	err = connection.Connect()
	if err != nil {
		return nil, err
	}

	rows, _, err := connection.Query("select * from pet")
	if err != nil {
		panic(err)
	}

	fmt.Printf("In table Pet %d rows.", len(rows))

	db.conn = connection

	return db, nil
}

func (db *Database) connect() (err error) {
	return nil
}
