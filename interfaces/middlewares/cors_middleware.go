package middlewares

import (
	"oees/application"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

type CORSMiddleware struct {
	Logger   hclog.Logger
	AppStore application.AppStore
}

func NewCORSMiddleware(logger hclog.Logger, appStore application.AppStore) *CORSMiddleware {
	return &CORSMiddleware{
		Logger:   logger,
		AppStore: appStore,
	}
}

func (cors *CORSMiddleware) AddCORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Vary", "Origin")
		ctx.Header("Vary", "Access-Control-Request-Method")
		ctx.Header("Vary", "Access-Control-Request-Headers")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH")
		ctx.Next()
	}
}
