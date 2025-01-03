package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	UpdateTaskByID(id uint, task Task) (Task, error)
	DeleteTaskByID(id uint) error
	GetTasksByUserID(userID uint) ([]Task, error) // Новый метод для задач по userID
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// (r *taskRepository) привязывает данную функцию к нашему репозиторию
func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}
func (r *taskRepository) GetTasksByUserID(userID uint) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	err := r.db.First(&task, &id).Error
	return task, err
}
func (r *taskRepository) DeleteTaskByID(id uint) error {
	return r.db.Delete(&Task{}, &id).Error
}
