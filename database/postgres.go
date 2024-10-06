package database

import (
	"fmt"

	"github.com/denaaay/task-management-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect(cred model.Cred) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cred.Host, cred.User, cred.Password, cred.DBName, cred.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
