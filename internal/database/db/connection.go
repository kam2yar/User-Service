package db

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type DBInterface interface {
	GetConnection() (*gorm.DB, error)
}

func DefaultConnection() *gorm.DB {
	var database DBInterface = Postgres{}
	connection, err := database.GetConnection()

	if err != nil {
		log.Panicln(fmt.Sprintf("database connenction abroted with error: %v", err))
	}

	return connection
}
