package model

type SignUpRequest struct {
	Nickname          string `json:"nickname" validate:"required"`
	Email             string `json:"email" validate:"required"`
	Password          string `json:"password" validate:"required"`
	VerificationCode  string `json:"verificationCode" validate:"required"`
	VerificationToken string `json:"verificationToken"`
}

func NewSignUpRequest() *SignUpRequest {
	return &SignUpRequest{}
}
