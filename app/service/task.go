package service

import (
	"go-todo/app/repository"
	"go-todo/app/request"
	"go-todo/app/response"
	"go-todo/pkg/logger"
)

type TaskService interface {
	List(req *request.TaskListRequest) (*response.TaskListResponse, error)
	Create(req *request.TaskCreateRequest) (*response.TaskResponse, error)
	Update(req *request.TaskUpdateRequest) (*response.TaskResponse, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *taskService {
	return &taskService{repo: repo}
}

func (service *taskService) List(req *request.TaskListRequest) (*response.TaskListResponse, error) {
	tasks, err := service.repo.List(req.GetLimit(), req.GetOffset(), req.Completed)
	if err != nil {
		logger.Log().Err(err).Msg("Failed to get tasks list")
		return nil, err
	}

	total, err := service.repo.Count(req.Completed)
	if err != nil {
		logger.Log().Err(err).Msg("Failed to get tasks count")
		return nil, err
	}

	return response.NewTaskListResponse(tasks, int(total), req.GetLimit(), req.GetOffset()), nil
}

func (service *taskService) Create(req *request.TaskCreateRequest) (*response.TaskResponse, error) {
	task, err := service.repo.Create(req.Title, req.Description)
	if err != nil {
		logger.Log().Err(err).Msg("Failed to create a task")
		return nil, err
	}

	return response.NewTaskResponse(task), nil
}

func (service *taskService) Update(req *request.TaskUpdateRequest) (*response.TaskResponse, error) {
	task, err := service.repo.Find(req.ID)

	if req.Title == "" {
		req.Title = task.Title
	}

	if req.Title == "" {
		req.Description = task.Description
	}

	if err != nil {
		logger.Log().Err(err).Msg("Failed to get a task")
		return nil, err
	}

	task, err = service.repo.Update(req.ID, req.Title, req.Description, req.Completed)
	if err != nil {
		logger.Log().Err(err).Msg("Failed to create a task")
		return nil, err
	}

	return response.NewTaskResponse(task), nil
}
