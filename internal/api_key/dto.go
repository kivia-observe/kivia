package apikey

type createApiKeyRequest struct {
	Name       string `json:"name"`
	ProjectId string `json:"project_id"`
}

type createApiKeyResponse struct {
	Message string `json:"message"`
	ApiKey string `json:"api_key"`
}