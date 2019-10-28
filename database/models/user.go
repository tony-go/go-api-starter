package models

import "github.com/jinzhu/gorm"

type Avatar struct {
	Type string
	Path string
	Size string
}

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Avatar Avatar `json:"avatar"`
	BirthDate string `json:"birthDate"`
	Password string `json:"password"`
}
