package controllers

type UserForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type ChangeEmailForm struct {
	Email string `json:"email" validate:"required,email"`
}

type VerificationEmailRequest struct {
	Body             string  `json:"body"`
	Recipient        string  `json:"recipient"`
	RecoveryCode     *string `json:"recovery_code"`
	RecoveryURL      *string `json:"recovery_url"`
	Subject          string  `json:"subject"`
	TemplateType     string  `json:"template_type"`
	To               string  `json:"to"`
	VerificationCode *string `json:"verification_code"`
	VerificationURL  *string `json:"verification_url"`
}

type LoginFlowAfterRequest struct {
	UserID          string    `json:"user_id"`
	Email           *string   `json:"email"`
	LoginChallenge  *string   `json:"oauth2_login_challenge"`
	IssuedAt        *string   `json:"issued_at"`
	RequestUrl      *string   `json:"request_url"`
	SecChUa         *[]string `json:"sec_ch_ua"`
	SecChUaMobile   *[]string `json:"sec_ch_ua_mobile"`
	SecChUaPlatform *[]string `json:"sec_ch_ua_platform"`
	UserAgent       *[]string `json:"user_agent"`
	XForwardedFor   *[]string `json:"x_forwarded_for"`
	RemoteAddr      *[]string `json:"remote_addr"`
}
