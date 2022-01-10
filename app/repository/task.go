package repository

import (
	"go-todo/app/model"
	"go-todo/app/repository/scope"
	"gorm.io/gorm"
)

var scopeWhereIfNotNil = scope.WhereIfNotNil

type TaskRepository interface {
	List(limit, offset int, completed *bool) ([]*model.Task, error)
	Count(completed *bool) (int64, error)
	Create(title, description string) (*model.Task, error)
	Update(id int, title, description string, completed bool) (*model.Task, error)
	Find(id int) (*model.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (repo *taskRepository) List(limit, offset int, completed *bool) (tasks []*model.Task, err error) {
	err = repo.db.Offset(offset).Limit(limit).Scopes(scopeWhereIfNotNil("completed", completed)).Find(&tasks).Error
	return
}

func (repo *taskRepository) Count(completed *bool) (count int64, err error) {
	err = repo.db.Model(&model.Task{}).Scopes(scopeWhereIfNotNil("completed", completed)).Count(&count).Error
	return
}

func (repo *taskRepository) Create(title, description string) (task *model.Task, err error) {
	task = &model.Task{
		Title:       title,
		Description: description,
	}
	err = repo.db.Model(&model.Task{}).Create(&task).Error
	return
}

func (repo *taskRepository) Update(id int, title, description string, completed bool) (task *model.Task, err error) {
	task = &model.Task{
		Title:       title,
		Description: description,
		Completed:   completed,
	}
	err = repo.db.Model(&model.Task{}).Where("id", id).Updates(&task).Error
	return
}

func (repo *taskRepository) Find(id int) (task *model.Task, err error) {
	err = repo.db.Model(&model.Task{}).Where("id", id).First(&task).Error
	return
}
