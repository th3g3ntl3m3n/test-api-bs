package api

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {

	DBHOST := os.Getenv("DBHOST")
	DBUSER := os.Getenv("DBUSER")
	DBPASS := os.Getenv("DBPASS")
	DBNAME := os.Getenv("DBNAME")
	DBPORT := os.Getenv("DBPORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", DBHOST, DBUSER, DBPASS, DBNAME, DBPORT)
	dbc, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Cannot open database: %v", err)
	}

	return dbc
}

func SetupTestDB() *gorm.DB {

	DBHOST := "localhost"
	DBUSER := "postgres"
	DBPASS := "root"
	DBNAME := "bondstate_db_test"
	DBPORT := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", DBHOST, DBUSER, DBPASS, DBNAME, DBPORT)
	dbc, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Cannot open database: %v", err)
	}

	return dbc
}
