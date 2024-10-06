package main

import (
	"fmt"
	"log"
	"os"

	"github.com/denaaay/task-management-api/database"
	"github.com/denaaay/task-management-api/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file : ", err)
	}

	var Credential = model.Cred{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}

	fmt.Println(Credential)

	db, err := database.DBConnect(Credential)
	if err != nil {
		log.Fatal("Error connect database : ", err)
	}

	errMigrate := db.AutoMigrate()
	if errMigrate != nil {
		log.Fatal("Error auto migrate : ", errMigrate)
	}

	fmt.Println("success connecting db : ", db)
}
