package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "xiaolizi"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = "3306"
	DATABASE = "mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	Conn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	// Conn := USERNAME + ":" + PASSWORD + "@" + NETWORK + "(" + SERVER + ":" + PORT + ")/" + DATABASE
	fmt.Println(Conn)
	SqlDB, err = sql.Open("mysql", Conn)
	//"root:xiaolizi@(localhost:3306)/mysql"

	if err != nil {
		log.Fatalln(err.Error())
	}

	err = SqlDB.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
