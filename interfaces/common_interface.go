package interfaces

import (
	"net/http"
	"oees/application"
	"oees/domain/value_objects"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

type CommonInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewCommonInterface(appStore *application.AppStore, logger hclog.Logger) *CommonInterface {
	return &CommonInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (commonInterface *CommonInterface) GetTables(ctx *gin.Context) {
	response := value_objects.Response{}

	tables, getErr := commonInterface.appStore.CommonApp.GetTables()
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Tables Found"
	response.Payload = tables

	ctx.JSON(http.StatusOK, response)
}
