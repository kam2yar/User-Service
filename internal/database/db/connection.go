package db

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DBInterface interface {
	GetConnection() (*gorm.DB, error)
}

func DefaultConnection() *gorm.DB {
	var database DBInterface = Postgres{}
	connection, err := database.GetConnection()

	if err != nil {
		zap.L().Panic(fmt.Sprintf("database connenction abroted with error: %v", err))
	}

	return connection
}
