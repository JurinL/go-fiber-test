package models

import "gorm.io/gorm"

// Person represents a person with a name and password.
type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

// User represents a user with a name, active status, and email.
type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

// Register represents a user registration with email, username, password, Line ID, phone, business kind, and website.
type Register struct {
	Email        string `json:"email" validate:"required,email,min=3,max=32"`
	Username     string `json:"username" validate:"required,min=3,max=32"`
	Password     string `json:"password" validate:"required,min=6,max=20"`
	LineID       string `json:"lineid" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	BusinessKind string `json:"businesskind" validate:"required"`
	Website      string `json:"website" validate:"required,url,min=2,max=30"`
}

type Dogs struct {
	gorm.Model 
	Name string `json:"name"`
	DogID int `json:"dog_id"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type ResultData struct {
	Data  []DogsRes `json:"data"`
	Name  string    `json:"name"`
	Count int       `json:"count"`
}

// Exercise 7.0.1
type Company struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	Email string `json:"email"`
	Website string `json:"website"`
	Facebook string `json:"facebook"`
}