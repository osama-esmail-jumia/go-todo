package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-todo/config"
	"go-todo/pkg/logger"
	"go-todo/router"
)

func main() {
	gin.SetMode(config.Cfg().AppMode)
	app := gin.Default()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	app.Use(cors.New(corsConfig))

	router.InitRouter(app)

	if err := app.Run(fmt.Sprintf(":%d", config.Cfg().AppPort)); err != nil {
		panic(err)
	}

	logger.Log().Printf("[info] start http server listening localhost:%d", config.Cfg().AppPort)
}
