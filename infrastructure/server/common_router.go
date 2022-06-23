package server

import (
	"oees/interfaces"
	"oees/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type CommonRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewCommonRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *CommonRouter {
	return &CommonRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (commonRouter *CommonRouter) ServeRoutes() {
	commonRouter.router.GET("/tables/", commonRouter.middlewares.AuthMiddleware.ValidateAccessToken(), commonRouter.interfaceStore.CommonInterface.GetTables)
}
