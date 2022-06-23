package middlewares

import (
	"oees/application"

	"github.com/hashicorp/go-hclog"
)

type MiddlewareStore struct {
	Logger         hclog.Logger
	AuthMiddleware *AuthMiddleware
	CORSMiddleware *CORSMiddleware
}

func NewMiddlewareStore(logger hclog.Logger, appStore *application.AppStore) *MiddlewareStore {
	middlewareStore := MiddlewareStore{}
	middlewareStore.Logger = logger
	middlewareStore.AuthMiddleware = NewAuthMiddleware(logger, *appStore)
	middlewareStore.CORSMiddleware = NewCORSMiddleware(logger, *appStore)
	return &middlewareStore
}
