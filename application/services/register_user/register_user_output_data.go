package services

import "time"

type RegisterUserOutputData struct {
	Email     string `validate:"required,email"`
	Password  string `validate:"required,min=8"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
