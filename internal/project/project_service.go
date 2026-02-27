package project

type projectservice struct {
	repo Repository
}

func NewProjectService(repo Repository) *projectservice {
	return &projectservice{
		repo: repo,
	}
}

func (s projectservice) CreateProject(project *Project) error {

	err := s.repo.Save(*project)

	if err != nil {
		return err
	}

	return nil
}