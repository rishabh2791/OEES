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

type SKUInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewskuInterface(appStore *application.AppStore, logger hclog.Logger) *SKUInterface {
	return &SKUInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (skuInterface *SKUInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	// Get new entry details from request body.
	model := entity.SKU{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create entry in database.
	created, creationErr := skuInterface.appStore.SKUApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "SKU Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (skuInterface *SKUInterface) CreateMultiple(ctx *gin.Context) {
	response := value_objects.Response{}
	createdModels := []interface{}{}
	creationErrors := []interface{}{}

	// Get new entry details from request body.
	models := []entity.SKU{}
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
		_, creationErr := skuInterface.appStore.SKUApp.Create(&model)
		if creationErr != nil {
			creationErrors = append(creationErrors, creationErr)
		} else {
			createdModels = append(createdModels, model)
		}
	}

	// Return response.
	response.Status = true
	response.Message = "SKU Created."
	response.Payload = map[string]interface{}{
		"models": createdModels,
		"errors": creationErrors,
	}

	ctx.JSON(http.StatusOK, response)
}

func (skuInterface *SKUInterface) Get(ctx *gin.Context) {
	response := value_objects.Response{}

	id := ctx.Param("id")
	sku, err := skuInterface.appStore.SKUApp.Get(id)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "SKU Found"
	response.Payload = sku

	ctx.JSON(http.StatusOK, response)
}

func (skuInterface *SKUInterface) List(ctx *gin.Context) {
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

	skus, err := skuInterface.appStore.SKUApp.List(utilities.ConvertJSONToSQL(conditions))
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "SKUs Found"
	response.Payload = skus

	ctx.JSON(http.StatusOK, response)
}

func (skuInterface *SKUInterface) Update(ctx *gin.Context) {
	response := value_objects.Response{}
	id := ctx.Param("id")

	// Get new entry details from request body.
	model := entity.SKU{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create entry in database.
	updated, creationErr := skuInterface.appStore.SKUApp.Update(id, &model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "SKU Updated."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}
