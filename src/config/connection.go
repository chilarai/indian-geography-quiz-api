package config

import (
	"config"
	"database/sql"
	"log"
)

// Mysqlconnect function
func Mysqlconnect() (mydb *sql.DB) {
	mydb, err := sql.Open(config.Mydriver, config.Myuser+":"+config.Mypass+"@/"+config.Mydb)
	if err != nil {
		log.Panic(err.Error())
	}
	return mydb
}