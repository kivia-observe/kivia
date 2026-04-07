package apikey

type createApiKeyRequest struct {
	Name       string `json:"name" validate:"required,min=3,max=15"`
	ProjectId string `json:"project_id" validate:"required,uuid"`
}

type createApiKeyResponse struct {
	Message string `json:"message"`
	ApiKey string `json:"api_key"`
}