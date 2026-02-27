package apikey

type createApiKeyRequest struct {
	Name       string `json:"name"`
	ProjectId string `json:"project_id"`
}