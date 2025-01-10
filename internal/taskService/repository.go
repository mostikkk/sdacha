package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	PostTask(task Task) (Task, error)
	PathTaskByID(id uint, task Task) (Task, error)
	DeleteTaskByID(id uint) error
	GetTasksByUserID(userID uint) ([]Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id uint, task *Task) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// (r *taskRepository) привязывает данную функцию к нашему репозиторию
func (r *taskRepository) PostTask(task Task) (Task, error) {
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

func (r *taskRepository) PathTaskByID(id uint, task Task) (Task, error) {
	err := r.db.First(&task, &id).Error
	return task, err
}
func (r *taskRepository) DeleteTaskByID(id uint) error {
	return r.db.Delete(&Task{}, &id).Error
}
func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}
func (r *taskRepository) GetTaskByID(id uint, task *Task) error {
	return r.db.First(task, id).Error
}
