package common

import (
	"database/sql"
	"log"
)

// Mysqlconnect function
func Mysqlconnect() (mydb *sql.DB) {
	mydb, err := sql.Open(Mydriver, Myuser+":"+Mypass+"@/"+Mydb)
	if err != nil {
		log.Panic(err.Error())
	}
	return mydb
}