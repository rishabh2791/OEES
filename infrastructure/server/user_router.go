package server

import (
	"oees/interfaces"
	"oees/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewUserRouter(router *gin.RouterGroup, interfaceStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *UserRouter {
	return &UserRouter{
		router:         router,
		interfaceStore: interfaceStore,
		middlewares:    middlewares,
	}
}

func (router *UserRouter) ServeRoutes() {
	router.router.POST("/create/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.UserInterface.Create)
	router.router.POST("/create/multi/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.UserInterface.CreateMultiple)
	router.router.GET("/:username/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.UserInterface.Get)
	router.router.POST("/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.UserInterface.List)
	router.router.PATCH("/:username/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.UserInterface.Update)
}
