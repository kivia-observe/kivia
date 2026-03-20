package user

type userservice struct {
	repo Repository
}
func NewUserService(repo Repository) *userservice {
	return &userservice{repo: repo}
}

func (s *userservice) GetUser(id string) (*User, error) {
	user := s.repo.FindById(id)
	
	if user.Id == "" {
		return nil, UserNotFound
	}
	
	return &user, nil
}

func (s *userservice) DeleteUser(id string) error {
	return s.repo.DeleteById(id)
}

func (s *userservice) UpdateUser(id string, user editUserRequest) error {
	return s.repo.UpdateById(id, user)
}