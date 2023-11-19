package change_email

import "time"

type ChangeEmailOutputData struct {
	Id        string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
