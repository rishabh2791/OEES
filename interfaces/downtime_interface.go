package interfaces

import (
	"encoding/json"
	"net/http"
	"oees/application"
	"oees/domain/entity"
	"oees/domain/value_objects"
	"oees/infrastructure/utilities"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

type DowntimeInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewDowntimeInterface(appStore *application.AppStore, logger hclog.Logger) *DowntimeInterface {
	return &DowntimeInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (downtimeInterface *DowntimeInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	// Get new entry details from request body.
	model := entity.Downtime{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if model.UpdatedByUsername == "" || len(model.UpdatedByUsername) == 0 {
		model.UpdatedByUsername = "rishabh2791"
	}

	// Create entry in database.
	created, creationErr := downtimeInterface.appStore.DowntimeApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "Downtime Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (downtimeInterface *DowntimeInterface) Get(ctx *gin.Context) {
	response := value_objects.Response{}

	id := ctx.Param("id")
	downtime, err := downtimeInterface.appStore.DowntimeApp.Get(id)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Downtime Found"
	response.Payload = downtime

	ctx.JSON(http.StatusOK, response)
}

func (downtimeInterface *DowntimeInterface) List(ctx *gin.Context) {
	response := value_objects.Response{}
	conditions := map[string]interface{}{}
	jsonError := json.NewDecoder(ctx.Request.Body).Decode(&conditions)
	if jsonError != nil {
		response.Status = false
		response.Message = jsonError.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	downtimes, err := downtimeInterface.appStore.DowntimeApp.List(utilities.ConvertJSONToSQL(conditions))
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Downtimes Found"
	response.Payload = downtimes

	ctx.JSON(http.StatusOK, response)
}

func (downtimeInterface *DowntimeInterface) Update(ctx *gin.Context) {
	response := value_objects.Response{}
	id := ctx.Param("id")

	// Get new entry details from request body.
	model := entity.Downtime{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if model.UpdatedByUsername == "" || len(model.UpdatedByUsername) == 0 {
		model.UpdatedByUsername = "rishabh2791"
	}

	// Create entry in database.
	updated, creationErr := downtimeInterface.appStore.DowntimeApp.Update(id, &model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "Downtime Updated."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}
