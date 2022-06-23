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

type ShiftInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewshiftInterface(appStore *application.AppStore, logger hclog.Logger) *ShiftInterface {
	return &ShiftInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (shiftInterface *ShiftInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	// Get new entry details from request body.
	model := entity.Shift{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create entry in database.
	created, creationErr := shiftInterface.appStore.ShiftApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "shift Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (shiftInterface *ShiftInterface) CreateMultiple(ctx *gin.Context) {
	response := value_objects.Response{}
	createdModels := []interface{}{}
	creationErrors := []interface{}{}

	// Get new entry details from request body.
	models := []entity.Shift{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&models)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = map[string]interface{}{
			"models": []string{},
			"errors": []string{jsonErr.Error()},
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	for _, model := range models {

		// Create entry in database.
		_, creationErr := shiftInterface.appStore.ShiftApp.Create(&model)
		if creationErr != nil {
			creationErrors = append(creationErrors, creationErr)
		} else {
			createdModels = append(createdModels, model)
		}
	}

	// Return response.
	response.Status = true
	response.Message = "Shift Created."
	response.Payload = map[string]interface{}{
		"models": createdModels,
		"errors": creationErrors,
	}

	ctx.JSON(http.StatusOK, response)
}

func (shiftInterface *ShiftInterface) Get(ctx *gin.Context) {
	response := value_objects.Response{}

	id := ctx.Param("id")
	shift, err := shiftInterface.appStore.ShiftApp.Get(id)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Shift Found"
	response.Payload = shift

	ctx.JSON(http.StatusOK, response)
}

func (shiftInterface *ShiftInterface) List(ctx *gin.Context) {
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

	shifts, err := shiftInterface.appStore.ShiftApp.List(utilities.ConvertJSONToSQL(conditions))
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Shifts Found"
	response.Payload = shifts

	ctx.JSON(http.StatusOK, response)
}

func (shiftInterface *ShiftInterface) Update(ctx *gin.Context) {
	response := value_objects.Response{}
	id := ctx.Param("id")

	// Get new entry details from request body.
	model := entity.Shift{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create entry in database.
	updated, creationErr := shiftInterface.appStore.ShiftApp.Update(id, &model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "Shift Updated."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}
