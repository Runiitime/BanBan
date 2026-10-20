package config

import (
	"fmt"
)

func DBConfig() string {
	dbName := "banban"
	dbUser := "postgres"
	dbPassword := "admin"
	dbHost := "localhost"
	dbPort := "5432"

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
}
