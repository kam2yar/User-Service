package main

import (
	"fmt"
	"github.com/kam2yar/user-service/internal/database/db"
	"github.com/kam2yar/user-service/internal/database/entities"
	"time"
)

var tables = []any{
	&entities.User{},
}

func main() {
	fmt.Println("Start migrating database structures", time.Now().Format(time.DateTime))
	migrate()
	fmt.Println("Migrations finished successfully", time.Now().Format(time.DateTime))
}

func migrate() {
	dbc := db.DefaultConnection()

	for index, table := range tables {
		fmt.Print(index + 1)
		fmt.Print(") Processing User... ")
		dbc.AutoMigrate(table)
		fmt.Println("(Done)")
	}
}
