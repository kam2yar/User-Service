package configurations

import (
	"github.com/kam2yar/user-service/internal/services"
	"strconv"
)

type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     int
}

var postgresPort, _ = strconv.Atoi(services.Env("POSTGRES_PORT"))
var Postgres DatabaseConfig = DatabaseConfig{
	Host:     services.Env("POSTGRES_HOST"),
	Username: services.Env("POSTGRES_USER"),
	Password: services.Env("POSTGRES_PASSWORD"),
	Name:     services.Env("POSTGRES_DB"),
	Port:     postgresPort,
}
