package router

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-todo/app/handler"
	"go-todo/app/repository"
	"go-todo/app/service"
	"go-todo/database"

	_ "go-todo/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter initialize routing information
// @title todo API
// @version 1.0
// @description todo API in go
// nolint
// @BasePath /api/v1
func InitRouter(r *gin.Engine) {
	taskRepository := repository.NewTaskRepository(database.DB())
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("/api/v1")
	apiV1.GET("/task", taskHandler.List)
	apiV1.POST("/task", taskHandler.Create)
	apiV1.PUT("/task/:id", taskHandler.Update)
}
