package apikey

import "time"

type ApiKey struct {

	Id string `json:"id"`

	Name string `json:"name"`

	Key string `json:"key"`

	UserId string `json:"user_id,omitempty"`

	ProjectId string `json:"project_id,omitempty"`

	Revoked bool `json:"revoked,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}