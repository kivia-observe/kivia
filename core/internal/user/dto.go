package user

type editUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}