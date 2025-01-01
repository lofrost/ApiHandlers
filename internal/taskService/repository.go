package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task) (Task, error)

	GetAllTasks() ([]Task, error)

	UpdateTaskByID(id uint, task Task) (Task, error)

	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	result := r.db.Model(&Task{}).Where("ID = ?", id).Updates(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	var UpdatedMessage Task
	r.db.Model(&Task{}).Where("ID = ?", id).First(&UpdatedMessage)
	return UpdatedMessage, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	result := r.db.Where("ID = ?", id).Delete(&Task{})
	return result.Error
}
