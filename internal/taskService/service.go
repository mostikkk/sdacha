package taskService

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}
func (s *TaskService) PostTask(task Task) (Task, error) {
	return s.repo.PostTask(task)
}

func (s *TaskService) GetTasksByUserID(userID uint) ([]Task, error) {
	return s.repo.GetTasksByUserID(userID)
}

func (s *TaskService) PathTask(id uint, task Task) (Task, error) {
	return s.repo.PathTaskByID(id, task)
}
func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
func (s *TaskService) GetAllTasks() ([]Task, error) {

	return s.repo.GetAllTasks()
}
func (s *TaskService) GetTaskByID(id uint) (Task, error) {
	var task Task
	err := s.repo.GetTaskByID(id, &task)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}
