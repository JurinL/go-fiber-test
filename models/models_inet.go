package models

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
	Website      string `json:"website" validate:"url"`
}

//,regexp=^[a-zA-Z0-9_]+$ regexp=^[0-9-]+$
