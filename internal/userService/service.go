package userService

type UserService struct {
	repo UserRepository
}

func NewService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}
func (s *UserService) GetUsers() ([]Users, error) {
	return s.repo.GetUsers()
}

func (s *UserService) PostUser(user Users) (Users, error) {
	return s.repo.PostUser(user)
}
func (s *UserService) PathUserByID(id uint, users Users) (Users, error) {
	return s.repo.PathUserByID(id, users)
}
func (s *UserService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}
