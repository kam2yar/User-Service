package configurations

import (
	"github.com/kam2yar/user-service/internal/helpers"
	"strconv"
)

type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     int
}

var postgresPort, _ = strconv.Atoi(helpers.Env("POSTGRES_PORT"))
var Postgres DatabaseConfig = DatabaseConfig{
	Host:     helpers.Env("POSTGRES_HOST"),
	Username: helpers.Env("POSTGRES_USER"),
	Password: helpers.Env("POSTGRES_PASSWORD"),
	Name:     helpers.Env("POSTGRES_DB"),
	Port:     postgresPort,
}
