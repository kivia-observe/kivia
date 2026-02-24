package refreshtoken

import "time"

type RefreshToken struct {
	Id string `json:"id"`

	Token string `json:"token"`

	Blacklisted bool `json:"blacklisted"`

	ExpiresAt time.Time `json:"expires_at"`

	UserId string `json:"user_id"`
}
