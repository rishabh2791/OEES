package server

import (
	"oees/interfaces"
	"oees/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type ShiftRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewShiftRouter(router *gin.RouterGroup, interfaceStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *ShiftRouter {
	return &ShiftRouter{
		router:         router,
		interfaceStore: interfaceStore,
		middlewares:    middlewares,
	}
}

func (router *ShiftRouter) ServeRoutes() {
	router.router.POST("/create/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.ShiftInterface.Create)
	router.router.POST("/create/multi/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.ShiftInterface.CreateMultiple)
	router.router.GET("/:id/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.ShiftInterface.Get)
	router.router.POST("/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.ShiftInterface.List)
	router.router.PATCH("/:id/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.ShiftInterface.Update)
}
