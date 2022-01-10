package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-todo/app/model"
	mock_repository "go-todo/app/repository/mock"
	"go-todo/app/request"
	"testing"
)

func TestTaskService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repository.NewMockTaskRepository(ctrl)
	service := NewTaskService(repo)
	req := request.TaskListRequest{}
	repo.EXPECT().List(0, 0, nil).Return([]*model.Task{
		&model.Task{
			Title:       "foo",
			Description: "bar",
		},
	}, nil)
	repo.EXPECT().Count(nil).Return(int64(5), nil)
	tasks, _ := service.List(&req)
	assert.Equal(t, 5, tasks.Total)
	assert.Equal(t, 1, len(tasks.Rows))
	assert.Equal(t, "foo", tasks.Rows[0].Title)
	assert.Equal(t, "bar", tasks.Rows[0].Description)
}

func TestTaskService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repository.NewMockTaskRepository(ctrl)
	service := NewTaskService(repo)
	req := request.TaskCreateRequest{
		Title:       "foo",
		Description: "bar",
	}
	repo.EXPECT().Create("foo", "bar").Return(&model.Task{
		Title:       "foo",
		Description: "bar",
	}, nil)
	task, _ := service.Create(&req)
	assert.Equal(t, "foo", task.Title)
	assert.Equal(t, "bar", task.Description)
}

func TestTaskService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repository.NewMockTaskRepository(ctrl)
	service := NewTaskService(repo)
	req := request.TaskUpdateRequest{
		ID:          1,
		Title:       "foo",
		Description: "bar",
		Completed:   true,
	}
	repo.EXPECT().Find(1).Return(&model.Task{
		Title:       "foo",
		Description: "bar",
		Completed:   true,
	}, nil)
	repo.EXPECT().Update(1, "foo", "bar", true).Return(&model.Task{
		Title:       "foo",
		Description: "bar",
		Completed:   true,
	}, nil)
	task, _ := service.Update(&req)
	assert.Equal(t, "foo", task.Title)
	assert.Equal(t, "bar", task.Description)
	assert.Equal(t, true, task.Completed)
}

func TestTaskService_List_ErrorWhenCallingList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repository.NewMockTaskRepository(ctrl)
	service := NewTaskService(repo)
	req := request.TaskListRequest{}
	repo.EXPECT().List(0, 0, nil).Return(nil, errors.New("error"))
	repo.EXPECT().Count(nil).Times(0)
	_, err := service.List(&req)
	assert.Error(t, err)
}

func TestTaskService_List_ErrorWhenCallingCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repository.NewMockTaskRepository(ctrl)
	service := NewTaskService(repo)
	req := request.TaskListRequest{}
	repo.EXPECT().List(0, 0, nil).Return([]*model.Task{
		&model.Task{
			Title:       "foo",
			Description: "bar",
		},
	}, nil)
	repo.EXPECT().Count(nil).Return(int64(0), errors.New("err"))
	_, err := service.List(&req)
	assert.Error(t, err)
}

func TestTaskService_Create_ErrorWhenCallingCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repository.NewMockTaskRepository(ctrl)
	service := NewTaskService(repo)
	req := request.TaskCreateRequest{
		Title:       "foo",
		Description: "bar",
	}
	repo.EXPECT().Create("foo", "bar").Return(nil, errors.New("err"))
	_, err := service.Create(&req)
	assert.Error(t, err)
}

func TestTaskService_Update_ErrorWhenCallingFind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repository.NewMockTaskRepository(ctrl)
	service := NewTaskService(repo)
	req := request.TaskUpdateRequest{
		ID:          1,
		Title:       "foo",
		Description: "bar",
		Completed:   true,
	}
	repo.EXPECT().Find(1).Return(nil, errors.New("err"))
	repo.EXPECT().Update(1, "foo", "bar", true).Times(0)
	_, err := service.Update(&req)
	assert.Error(t, err)
}

func TestTaskService_Update_ErrorWhenCallingUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_repository.NewMockTaskRepository(ctrl)
	service := NewTaskService(repo)
	req := request.TaskUpdateRequest{
		ID:          1,
		Title:       "foo",
		Description: "bar",
		Completed:   true,
	}
	repo.EXPECT().Find(1).Return(&model.Task{
		Title:       "foo",
		Description: "bar",
		Completed:   true,
	}, nil)
	repo.EXPECT().Update(1, "foo", "bar", true).Return(nil, errors.New("err"))
	_, err := service.Update(&req)
	assert.Error(t, err)
}
