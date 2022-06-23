package server

import (
	"oees/interfaces"
	"oees/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type PlantRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewPlantRouter(router *gin.RouterGroup, interfaceStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *PlantRouter {
	return &PlantRouter{
		router:         router,
		interfaceStore: interfaceStore,
		middlewares:    middlewares,
	}
}

func (router *PlantRouter) ServeRoutes() {
	router.router.POST("/create/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.PlantInterface.Create)
	router.router.POST("/create/multi/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.PlantInterface.CreateMultiple)
	router.router.GET("/:id/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.PlantInterface.Get)
	router.router.POST("/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.PlantInterface.List)
	router.router.PATCH("/:id/", router.middlewares.CORSMiddleware.AddCORSMiddleware(), router.middlewares.AuthMiddleware.ValidateAccessToken(), router.interfaceStore.PlantInterface.Update)
}
