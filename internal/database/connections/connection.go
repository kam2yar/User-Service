package connections

import (
	"fmt"
	"gorm.io/gorm"
)

type DBInterface interface {
	GetConnection() (*gorm.DB, error)
}

func DefaultConnection() *gorm.DB {
	var database DBInterface = Postgres{}
	connection, err := database.GetConnection()

	if err != nil {
		panic(fmt.Sprintf("Database connenction abroted with error: %v", err))
	}

	return connection
}
