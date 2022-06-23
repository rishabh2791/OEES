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

type UserRoleInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewuserRoleInterface(appStore *application.AppStore, logger hclog.Logger) *UserRoleInterface {
	return &UserRoleInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (userRoleInterface *UserRoleInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	// Get new entry details from request body.
	model := entity.UserRole{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create entry in database.
	created, creationErr := userRoleInterface.appStore.UserRoleApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "User Role Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (userRoleInterface *UserRoleInterface) CreateMultiple(ctx *gin.Context) {
	response := value_objects.Response{}
	createdModels := []interface{}{}
	creationErrors := []interface{}{}

	// Get new entry details from request body.
	models := []entity.UserRole{}
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
		created, creationErr := userRoleInterface.appStore.UserRoleApp.Create(&model)
		if creationErr != nil {
			creationErrors = append(creationErrors, creationErr)
		} else {
			createdModels = append(createdModels, created)
		}
	}

	// Return response.
	response.Status = true
	response.Message = "User Role Created."
	response.Payload = map[string]interface{}{
		"models": createdModels,
		"errors": creationErrors,
	}

	ctx.JSON(http.StatusOK, response)
}

func (userRoleInterface *UserRoleInterface) List(ctx *gin.Context) {
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

	userRoles, err := userRoleInterface.appStore.UserRoleApp.List(utilities.ConvertJSONToSQL(conditions))
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "User Roles Found"
	response.Payload = userRoles

	ctx.JSON(http.StatusOK, response)
}
