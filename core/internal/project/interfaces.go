package project

type ProjectRepository interface {
	Save(project Project) error
	FindById(id string) (*Project, error)
	ExistsById(id string) (bool, error)
	FindProjectIdByApiKey(apiKey string) (string, error)
	FindAllByUserId(userId string) ([]ProjectResponse, error)
}
