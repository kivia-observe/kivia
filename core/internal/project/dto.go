package project

import "time"

type createProjectRequest struct {
	Name string `json:"name" validate:"required,min=3,max=15"`
}

type ProjectResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	ApiKeys   int       `json:"api_keys"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}