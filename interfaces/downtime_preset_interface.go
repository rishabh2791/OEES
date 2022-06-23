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

type DowntimePresetInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewDowntimePresetInterface(appStore *application.AppStore, logger hclog.Logger) *DowntimePresetInterface {
	return &DowntimePresetInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (downtimePresetInterface *DowntimePresetInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	// Get new entry details from request body.
	model := entity.PresetDowntime{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create entry in database.
	created, creationErr := downtimePresetInterface.appStore.DowntimePresetApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "Preset Downtime Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (downtimePresetInterface *DowntimePresetInterface) CreateMultiple(ctx *gin.Context) {
	response := value_objects.Response{}
	createdModels := []interface{}{}
	creationErrors := []interface{}{}

	// Get new entry details from request body.
	models := []entity.PresetDowntime{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&models)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	for _, model := range models {

		// Create entry in database.
		created, creationErr := downtimePresetInterface.appStore.DowntimePresetApp.Create(&model)
		if creationErr != nil {
			creationErrors = append(creationErrors, creationErr)
		} else {
			createdModels = append(createdModels, created)
		}
	}

	// Return response.
	response.Status = true
	response.Message = "Preset Downtime Created."
	response.Payload = map[string]interface{}{
		"models": createdModels,
		"errors": creationErrors,
	}

	ctx.JSON(http.StatusOK, response)
}

func (downtimePresetInterface *DowntimePresetInterface) List(ctx *gin.Context) {
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

	devices, err := downtimePresetInterface.appStore.DowntimePresetApp.List(utilities.ConvertJSONToSQL(conditions))
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Preset Downtimes Found"
	response.Payload = devices

	ctx.JSON(http.StatusOK, response)
}
