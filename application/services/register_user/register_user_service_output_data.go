package services

import "time"

type RegisterUserOutputData struct {
	Id        string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
