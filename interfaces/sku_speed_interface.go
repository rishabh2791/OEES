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

type SKUSpeedInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewskuSpeedInterface(appStore *application.AppStore, logger hclog.Logger) *SKUSpeedInterface {
	return &SKUSpeedInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (skuSpeedInterface *SKUSpeedInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	// Get new entry details from request body.
	model := entity.SKUSpeed{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create entry in database.
	created, creationErr := skuSpeedInterface.appStore.SKUSpeedApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "SKU Speed Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (skuSpeedInterface *SKUSpeedInterface) CreateMultiple(ctx *gin.Context) {
	response := value_objects.Response{}
	createdModels := []interface{}{}
	creationErrors := []interface{}{}

	// Get new entry details from request body.
	models := []entity.SKUSpeed{}
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
		_, creationErr := skuSpeedInterface.appStore.SKUSpeedApp.Create(&model)
		if creationErr != nil {
			creationErrors = append(creationErrors, creationErr)
		} else {
			createdModels = append(createdModels, model)
		}
	}

	// Return response.
	response.Status = true
	response.Message = "SKU Speed Created."
	response.Payload = map[string]interface{}{
		"models": createdModels,
		"errors": creationErrors,
	}

	ctx.JSON(http.StatusOK, response)
}

func (skuSpeedInterface *SKUSpeedInterface) Get(ctx *gin.Context) {
	response := value_objects.Response{}

	id := ctx.Param("id")
	skuSpeed, err := skuSpeedInterface.appStore.SKUSpeedApp.Get(id)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "SKU Speed Found"
	response.Payload = skuSpeed

	ctx.JSON(http.StatusOK, response)
}

func (skuSpeedInterface *SKUSpeedInterface) List(ctx *gin.Context) {
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

	skuSpeeds, err := skuSpeedInterface.appStore.SKUSpeedApp.List(utilities.ConvertJSONToSQL(conditions))
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "SKU Speeds Found"
	response.Payload = skuSpeeds

	ctx.JSON(http.StatusOK, response)
}

func (skuSpeedInterface *SKUSpeedInterface) Update(ctx *gin.Context) {
	response := value_objects.Response{}
	id := ctx.Param("id")

	// Get new entry details from request body.
	model := entity.SKUSpeed{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create entry in database.
	updated, creationErr := skuSpeedInterface.appStore.SKUSpeedApp.Update(id, &model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "SKU Speed Updated."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}
