package emailverification

import "time"

type EmailVerification struct {
	Id        string
	Email     string
	OtpHash   string
	Attempts  int
	ExpiresAt time.Time
	CreatedAt time.Time
	UserId    string
}
