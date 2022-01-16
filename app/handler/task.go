package handler

import (
	"github.com/gin-gonic/gin"
	"go-todo/app/request"
	_ "go-todo/app/request"
	"go-todo/app/response"
	"go-todo/app/service"
	"net/http"
	"strings"
)

type TaskHandler interface {
	List(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
}

func NewTaskHandler(service service.TaskService) TaskHandler {
	return &taskHandler{service: service}
}

type taskHandler struct {
	service service.TaskService
}

// List GetTasks @Summary Get tasks list
// @Router /task [get]
// @Tags task
// @Produce json
// @Param limit query int false "pagination limit"
// @Param offset query int false "pagination offset"
// @Param completed query boolean false "task status"
// @Success 200 {array} response.TaskListResponse
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
func (handler *taskHandler) List(context *gin.Context) {
	req, err := request.NewTaskListRequest(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.NewValidationErrorResponse(err))
		return
	}

	resp, err := handler.service.List(&req)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.NewBadRequestResponse())
		return
	}

	context.JSON(http.StatusOK, resp)
}

// Create CreateTask @Summary create a new task
// @Router /task [post]
// @Tags task
// @Accept json
// @Produce json
// @Param body body request.TaskCreateRequest true "task body"
// @Success 200 {array} response.TaskResponse
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
func (handler *taskHandler) Create(context *gin.Context) {
	req, err := request.NewTaskCreateRequest(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.NewValidationErrorResponse(err))
		return
	}

	resp, err := handler.service.Create(&req)
	if err != nil {
		if strings.Contains(err.Error(), service.DUPLICATE_ERROR) {
			context.JSON(http.StatusBadRequest, response.NewErrorResponse(response.DUPLICATE_TITLE_MESSAGE))
			return
		}
		context.JSON(http.StatusBadRequest, response.NewBadRequestResponse())
		return
	}

	context.JSON(http.StatusCreated, resp)
}

// Update CreateTask @Summary create a new task
// @Router /task/{id} [put]
// @Tags task
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param body body request.TaskUpdateRequest true "task body"
// @Success 200 {array} response.TaskResponse
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
func (handler *taskHandler) Update(context *gin.Context) {
	req, err := request.NewTaskUpdateRequest(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.NewValidationErrorResponse(err))
		return
	}

	resp, err := handler.service.Update(&req)
	if err != nil {
		if err.Error() == service.NOT_FOUND_ERROR {
			context.JSON(http.StatusNotFound, response.NewNotFoundResponse())
			return
		}
		if strings.Contains(err.Error(), service.DUPLICATE_ERROR) {
			context.JSON(http.StatusBadRequest, response.NewErrorResponse(response.DUPLICATE_TITLE_MESSAGE))
			return
		}
		context.JSON(http.StatusBadRequest, response.NewBadRequestResponse())
		return
	}

	context.JSON(http.StatusOK, resp)
}
