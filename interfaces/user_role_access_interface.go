package interfaces

import (
	"encoding/json"
	"log"
	"net/http"
	"oees/application"
	"oees/domain/entity"
	"oees/domain/value_objects"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

type UserRoleAccessInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewUserRoleAccessInterface(appStore *application.AppStore, logger hclog.Logger) *UserRoleAccessInterface {
	return &UserRoleAccessInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (userRoleAccessInterface *UserRoleAccessInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	requestingUser, ok := ctx.Get("user")
	if !ok {
		response.Status = false
		response.Message = "Anonymous User"
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	user := requestingUser.(*entity.User)

	model := entity.UserRoleAccess{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	model.CreatedByUsername = user.Username
	model.UpdatedByUsername = user.Username

	created, creationErr := userRoleAccessInterface.appStore.UserRoleAccessApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "User Role Access Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (userRoleAccessInterface *UserRoleAccessInterface) CreateMultiple(ctx *gin.Context) {
	response := value_objects.Response{}
	createdModels := []interface{}{}
	creationErrors := []interface{}{}

	requestingUser, ok := ctx.Get("user")
	if !ok {
		response.Status = false
		response.Message = "Anonymous User"
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	user := requestingUser.(*entity.User)

	models := []entity.UserRoleAccess{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&models)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	for _, model := range models {
		model.CreatedByUsername = user.Username
		model.UpdatedByUsername = user.Username

		_, creationErr := userRoleAccessInterface.appStore.UserRoleAccessApp.Create(&model)
		if creationErr != nil {
			log.Println(creationErr)
			creationErrors = append(creationErrors, creationErr)
		} else {
			createdModels = append(createdModels, model)
		}
	}

	response.Status = true
	response.Message = "User Role Accesses Created."
	response.Payload = map[string]interface{}{
		"models": createdModels,
		"errors": creationErrors,
	}

	ctx.JSON(http.StatusOK, response)
}

func (userRoleAccessInterface *UserRoleAccessInterface) List(ctx *gin.Context) {
	response := value_objects.Response{}
	userRole := ctx.Param("user_role")

	accesses, getErr := userRoleAccessInterface.appStore.UserRoleAccessApp.List(userRole)
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "User Role Accesses Found"
	response.Payload = accesses

	ctx.JSON(http.StatusOK, response)
}

func (userRoleAccessInterface *UserRoleAccessInterface) Update(ctx *gin.Context) {
	response := value_objects.Response{}
	userRole := ctx.Param("user_role")

	requestingUser, ok := ctx.Get("user")
	if !ok {
		response.Status = false
		response.Message = "Anonymous User."
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	user := requestingUser.(*entity.User)

	model := entity.UserRoleAccess{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	model.UpdatedByUsername = user.Username

	updated, updationErr := userRoleAccessInterface.appStore.UserRoleAccessApp.Update(userRole, &model)
	if updationErr != nil {
		response.Status = false
		response.Message = updationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "User Role Access Updated."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}
