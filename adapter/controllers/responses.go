package controllers

import "time"

type userResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newUserResponse(id string, email string, created_at time.Time, updated_at time.Time) *userResponse {
	return &userResponse{
		ID:        id,
		Email:     email,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}
}
