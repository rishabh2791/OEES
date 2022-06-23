package server

import (
	"oees/interfaces"
	"oees/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type LineRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewLineRouter(router *gin.RouterGroup, interfaceStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *LineRouter {
	return &LineRouter{
		router:         router,
		interfaceStore: interfaceStore,
		middlewares:    middlewares,
	}
}

func (router *LineRouter) ServeRoutes() {
	router.router.POST("/create/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.LineInterface.Create)
	router.router.POST("/create/multi/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.LineInterface.CreateMultiple)
	router.router.GET("/:id/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.LineInterface.Get)
	router.router.POST("/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.LineInterface.List)
	router.router.PATCH("/:id/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.LineInterface.Update)
}
