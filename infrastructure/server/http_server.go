package server

import (
	"io"
	"net/http"
	"oees/application"
	"oees/domain/value_objects"
	"oees/infrastructure/utilities"
	"oees/interfaces"
	"oees/interfaces/middlewares"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HTTPServer struct {
	Router          *gin.Engine
	MiddlewareStore *middlewares.MiddlewareStore
	InterfaceStore  *interfaces.InterfaceStore
	AppStore        *application.AppStore
}

func NewHTTPServer(serverConfig utilities.ServerConfig, appStore *application.AppStore, interfaceStore *interfaces.InterfaceStore, middlewareStore *middlewares.MiddlewareStore) *HTTPServer {
	httpServer := HTTPServer{}

	if !serverConfig.IsDebug() {
		gin.SetMode(gin.ReleaseMode)
	}

	utilities.GinLogger() // Comment out for terminal logging
	httpServer.Router = gin.Default()
	httpServer.Router.Static("/public", "/home/pi/Development/backend/public")
	httpServer.InterfaceStore = interfaceStore
	httpServer.MiddlewareStore = middlewareStore
	httpServer.AppStore = appStore

	return &httpServer
}

func (httpServer *HTTPServer) Serve() {
	httpServer.Router.POST("/image/upload/", ImageUploader)
	authRouter := NewAuthRouter(httpServer.Router.Group("/auth/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	commonRouter := NewCommonRouter(httpServer.Router.Group("/common/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	deviceRouter := NewDeviceRouter(httpServer.Router.Group("/device/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	deviceDataRouter := NewDeviceDataRouter(httpServer.Router.Group("/device_data/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	downtimeRouter := NewDowntimeRouter(httpServer.Router.Group("/downtime/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	downtimePresetRouter := NewDowntimePresetRouter(httpServer.Router.Group("/preset/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	jobRouter := NewJobRouter(httpServer.Router.Group("/job/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	lineRouter := NewLineRouter(httpServer.Router.Group("/line/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	shiftRouter := NewShiftRouter(httpServer.Router.Group("/shift/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	skuRouter := NewSKURouter(httpServer.Router.Group("/sku/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	taskRouter := NewTaskRouter(httpServer.Router.Group("/task/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	taskBatchRouter := newTaskBatchRouter(httpServer.Router.Group("/task_batch/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	userRouter := NewUserRouter(httpServer.Router.Group("/user/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	userRoleRouter := NewUserRoleRouter(httpServer.Router.Group("/user_role/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	userRoleAccessRouter := NewUserRoleAccessRouter(httpServer.Router.Group("/user_role_access/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)

	authRouter.ServeRoutes()
	commonRouter.ServeRoutes()
	deviceDataRouter.ServeRoutes()
	deviceRouter.ServeRoutes()
	downtimeRouter.ServeRoutes()
	downtimePresetRouter.ServeRoutes()
	jobRouter.ServeRoutes()
	lineRouter.ServeRoutes()
	shiftRouter.ServeRoutes()
	skuRouter.ServeRoutes()
	taskRouter.ServeRoutes()
	taskBatchRouter.ServeRoutes()
	userRoleRouter.ServeRoutes()
	userRouter.ServeRoutes()
	userRoleAccessRouter.ServeRoutes()
}

func ImageUploader(ctx *gin.Context) {
	response := value_objects.Response{}

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	extension := strings.Split(header.Filename, ".")[len(strings.Split(header.Filename, "."))-1]
	filename := "/home/pi/Development/backend/public/profile_pics/" + uuid.New().String() + "." + extension
	out, err := os.Create(filename)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Image Uploaded"
	response.Payload = filename

	ctx.AbortWithStatusJSON(http.StatusOK, response)
	return
}
