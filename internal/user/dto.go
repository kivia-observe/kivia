package user

type createUserRequest struct {

	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password,omitempty`

}

