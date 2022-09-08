package server

import (
	"oees/interfaces"
	"oees/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRoleAccessRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewUserRoleAccessRouter(router *gin.RouterGroup, interfaceStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *UserRoleAccessRouter {
	return &UserRoleAccessRouter{
		router:         router,
		interfaceStore: interfaceStore,
		middlewares:    middlewares,
	}
}

func (userRoleAccessRouter *UserRoleAccessRouter) ServeRoutes() {
	userRoleAccessRouter.router.POST("/create/", userRoleAccessRouter.middlewares.AuthMiddleware.ValidateAccessToken(), userRoleAccessRouter.interfaceStore.UserRoleAccessInterface.Create)
	userRoleAccessRouter.router.POST("/create/multi/", userRoleAccessRouter.middlewares.AuthMiddleware.ValidateAccessToken(), userRoleAccessRouter.interfaceStore.UserRoleAccessInterface.CreateMultiple)
	userRoleAccessRouter.router.GET("/:user_role/", userRoleAccessRouter.middlewares.AuthMiddleware.ValidateAccessToken(), userRoleAccessRouter.interfaceStore.UserRoleAccessInterface.List)
	userRoleAccessRouter.router.PATCH("/:user_role/", userRoleAccessRouter.middlewares.AuthMiddleware.ValidateAccessToken(), userRoleAccessRouter.interfaceStore.UserRoleAccessInterface.Update)
}
