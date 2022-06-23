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

type UserInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewUserInterface(appStore *application.AppStore, logging hclog.Logger) *UserInterface {
	return &UserInterface{
		appStore: appStore,
		logger:   logging,
	}
}

func (userInterface *UserInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	// Get new entry details from request body.
	model := entity.User{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create entry in database.
	created, creationErr := userInterface.appStore.UserApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "User Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (userInterface *UserInterface) CreateMultiple(ctx *gin.Context) {
	response := value_objects.Response{}
	createdModels := []interface{}{}
	creationErrors := []interface{}{}

	// Get new entry details from request body.
	models := []entity.User{}
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
		_, creationErr := userInterface.appStore.UserApp.Create(&model)
		if creationErr != nil {
			creationErrors = append(creationErrors, creationErr)
		} else {
			createdModels = append(createdModels, model)
		}
	}

	// Return response.
	response.Status = true
	response.Message = "User Created."
	response.Payload = map[string]interface{}{
		"models": createdModels,
		"errors": creationErrors,
	}

	ctx.JSON(http.StatusOK, response)
}

func (userInterface *UserInterface) Get(ctx *gin.Context) {
	response := value_objects.Response{}

	username := ctx.Param("username")
	user, err := userInterface.appStore.UserApp.Get(username)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "User Found"
	response.Payload = user

	ctx.JSON(http.StatusOK, response)
}

func (userInterface *UserInterface) List(ctx *gin.Context) {
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

	users, err := userInterface.appStore.UserApp.List(utilities.ConvertJSONToSQL(conditions))
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Users Found"
	response.Payload = users

	ctx.JSON(http.StatusOK, response)
}

func (userInterface *UserInterface) Update(ctx *gin.Context) {
	response := value_objects.Response{}
	username := ctx.Param("username")

	// Get new entry details from request body.
	model := entity.User{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Complete model details
	if len(model.Password) != 0 {
		hashedPass, err := utilities.Hash(model.Password)
		if err != nil {
		} else {
			model.Password = string(hashedPass)
		}
	}

	// Create entry in database.
	updated, creationErr := userInterface.appStore.UserApp.Update(username, &model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "User Updated."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}
