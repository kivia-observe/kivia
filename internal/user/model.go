package user

import "time"

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"-"`
	JoinedAt time.Time `json:"joined_at"`
}