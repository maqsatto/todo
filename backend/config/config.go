package config

import (
	"os"
)

var JwtKey = []byte("BAUKA_GOI")

func GetDSN() string {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=0000 dbname=todolist port=5432 sslmode=disable"
	}
	return dsn
}
