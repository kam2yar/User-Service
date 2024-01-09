package db

import (
	"fmt"
	"github.com/kam2yar/user-service/internal/configurations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct{}

func (Postgres) GetConnection() (*gorm.DB, error) {
	var config = configurations.Postgres

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tehran",
			config.Host, config.Username, config.Password, config.Name, config.Port,
		),
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	return db, err
}
