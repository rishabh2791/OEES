package server

import (
	"oees/interfaces"
	"oees/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewAuthRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *AuthRouter {
	return &AuthRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (authRouter *AuthRouter) ServeRoutes() {
	authRouter.router.POST("/login/", authRouter.middlewares.CORSMiddleware.AddCORSMiddleware(), authRouter.middlewares.AuthMiddleware.ValidateUser(), authRouter.interfaceStore.AuthInterface.Login)
	authRouter.router.POST("/logout/", authRouter.middlewares.CORSMiddleware.AddCORSMiddleware(), authRouter.middlewares.AuthMiddleware.ValidateAccessToken(), authRouter.interfaceStore.AuthInterface.Logout)
	authRouter.router.GET("/refresh/", authRouter.middlewares.CORSMiddleware.AddCORSMiddleware(), authRouter.middlewares.AuthMiddleware.ValidateRefreshToken(), authRouter.interfaceStore.AuthInterface.Refresh)
	authRouter.router.POST("/reset/password/", authRouter.middlewares.CORSMiddleware.AddCORSMiddleware(), authRouter.middlewares.AuthMiddleware.ValidateAccessToken(), authRouter.interfaceStore.AuthInterface.ResetPassword)
}
