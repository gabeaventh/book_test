package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // import the PostgreSQL driver
	"github.com/spf13/viper"
)

func InitDB() *sql.DB {
	viper.SetConfigFile("config.toml")
	viper.ReadInConfig()
	dsn := viper.GetString("database.dsn")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	return db
}
