package auth

type createUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
