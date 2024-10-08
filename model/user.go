package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Email    string `gorm:"type:varhcar(100);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
}

type CreateUserResp struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
