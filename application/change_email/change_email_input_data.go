package change_email

type ChangeEmailInputData struct {
	Id    string `json:"id" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
