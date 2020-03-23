package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
)

var connection *gorm.DB

func Initialize() *gorm.DB {
	connString := ConnectionString()
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func ConnectionString() string {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USERNAME")
	dbName := os.Getenv("DB_DATABASE")

	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbName,
	)
}

func GetConnection() *gorm.DB {
	if connection == nil {
		connection = Initialize()
	}

	return connection
}