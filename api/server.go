package server

import (
	"github.com/gin-gonic/gin"

	controller "golang-gin-api/api/controller"
	middleware "golang-gin-api/api/middleware"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(tagController *controller.TagController) *ServerHTTP {
	engine := gin.New()

	engine.ForwardedByClientIP = true
	engine.SetTrustedProxies([]string{"localhost"})

	engine.Use(gin.Logger()) // use logger from gin

	engine.POST("/login", middleware.LoginHandler) // request jwt

	api := engine.Group("/api", middleware.AuthorizationMiddleware)

	api.GET("tags", tagController.FindAll)
	api.GET("tags/:id", tagController.FindByID)
	api.POST("tags", tagController.Save)
	api.DELETE("tags/:id", tagController.Delete)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8080")
}
