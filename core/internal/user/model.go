package user

import "time"

type User struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Password       string    `json:"-"`
	Role           string    `json:"role"`
	ProfilePicture string    `json:"profile_picture"`
	Provider       string    `json:"provider"`
	Verified       bool      `json:"verified"`
	JoinedAt       time.Time `json:"joined_at"`
}
