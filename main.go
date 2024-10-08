package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/denaaay/task-management-api/database"
	"github.com/denaaay/task-management-api/model"
	"github.com/denaaay/task-management-api/router"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file : ", err)
	}

	Credential := model.Cred{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}

	db, err := database.DBConnect(Credential)
	if err != nil {
		log.Fatal("Error connect database : ", err)
	}

	errMigrate := db.AutoMigrate(&model.User{})
	if errMigrate != nil {
		log.Fatal("Error auto migrate : ", errMigrate)
	}

	fmt.Println("success connecting db : ", db)
	server(db)
}

func server(_ *gorm.DB) {

	server := &http.Server{
		Addr:    ":3000",
		Handler: router.NewRouter(),
	}

	server.ListenAndServe()
}
