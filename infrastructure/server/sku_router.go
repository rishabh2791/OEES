package server

import (
	"oees/interfaces"
	"oees/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type SKURouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewSKURouter(router *gin.RouterGroup, interfaceStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *SKURouter {
	return &SKURouter{
		router:         router,
		interfaceStore: interfaceStore,
		middlewares:    middlewares,
	}
}

func (router *SKURouter) ServeRoutes() {
	router.router.POST("/create/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.SKUInterface.Create)
	router.router.POST("/create/multi/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.SKUInterface.CreateMultiple)
	router.router.GET("/:id/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.SKUInterface.Get)
	router.router.POST("/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.SKUInterface.List)
	router.router.PATCH("/:id/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.SKUInterface.Update)
}
