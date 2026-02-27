package apikey

import "time"

type ApiKey struct {

	Id string

	Name string

	Key string

	UserId string

	ProjectId string 

	Revoked bool

	CreatedAt time.Time
}