package interfaces

import (
	"net/http"
	"oees/application"
	"oees/domain/entity"
	"oees/domain/value_objects"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

type AuthInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewAuthInterface(appStore *application.AppStore, logging hclog.Logger) *AuthInterface {
	return &AuthInterface{
		appStore: appStore,
		logger:   logging,
	}
}

func (auth *AuthInterface) Login(ctx *gin.Context) {
	response := value_objects.Response{}

	requestingUser, ok := ctx.Get("user")
	if !ok {
		response.Status = false
		response.Message = "User Invalid."
		response.Payload = "User in Context."
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Get User details about the requestig user from the repository
	user, err := auth.appStore.UserApp.Get(requestingUser.(*entity.User).Username)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = "User Details."
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Validate the requesting user and repository user by comparing their passwords.
	validationErr := auth.appStore.AuthApp.Authenticate(requestingUser.(*entity.User), user)
	if validationErr != nil {
		response.Status = false
		response.Message = validationErr.Error()
		response.Payload = "User Authentication"
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// On Successful authentication generate pair of access and refresh tokens.
	tokens, tokenGenErr := auth.appStore.AuthApp.GenerateTokens(user)
	if tokenGenErr != nil {
		response.Status = false
		response.Message = tokenGenErr.Error()
		response.Payload = "Token Generation"
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Store Access and Refresh UUIDs in Redis Server
	tokenErr := auth.appStore.AuthApp.GenerateAuth(tokens)
	if tokenErr != nil {
		response.Status = false
		response.Message = tokenErr.Error()
		response.Payload = "Storing Token."
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Login Successful."
	response.Payload = tokens
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "POST")
	ctx.JSON(http.StatusOK, response)
}

func (auth *AuthInterface) Logout(ctx *gin.Context) {
	response := value_objects.Response{}

	// Check User presence in Request Context
	user, ok := ctx.Get("user")
	if !ok {
		response.Status = false
		response.Message = "User Invalid."
		response.Payload = "User in Context."
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Check AccessDetails presence in Request Context
	accessDetails, ok := ctx.Get("access_details")
	if !ok {
		response.Status = false
		response.Message = "Access Details Invalid."
		response.Payload = "Access Details in Context."
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Delete Access UUID from Redis Server
	_, err := auth.appStore.AuthApp.DeleteAuth(accessDetails.(*value_objects.AccessDetail).AccessUUID)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = "Unable to Logout."
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "User Logged Out."
	response.Payload = user.(*entity.User)
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, response)
}

func (auth *AuthInterface) Refresh(ctx *gin.Context) {
	response := value_objects.Response{}

	// Check User presence in Request Context
	user, ok := ctx.Get("user")
	if !ok {
		response.Status = false
		response.Message = "User Invalid."
		response.Payload = "User in Context."
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Check Refresh Details presence in Request Context
	refreshDetails, ok := ctx.Get("refresh_details")
	if !ok {
		response.Status = false
		response.Message = "Access Details Invalid."
		response.Payload = "Access Details in Context."
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// On presence of User and Refresh Details in Request Context, delete old refresh UUID from Redis Server
	_, err := auth.appStore.AuthApp.DeleteAuth(refreshDetails.(*value_objects.RefreshDetail).RefreshUUID)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = "Unable to Delete Old Authentication Details."
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		// On Successful deletion of refresh UUID from Redis Server, generate new pair of access and refresh tokens.
		tokens, tokenGenErr := auth.appStore.AuthApp.GenerateTokens(user.(*entity.User))
		if tokenGenErr != nil {
			response.Status = false
			response.Message = tokenGenErr.Error()
			response.Payload = "Token Generation"
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		// On Successful generation of Tokens store Access and Refresh UUIDs in Redis Server.
		tokenErr := auth.appStore.AuthApp.GenerateAuth(tokens)
		if tokenErr != nil {
			response.Status = false
			response.Message = tokenErr.Error()
			response.Payload = "Storing Token."
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		response.Status = true
		response.Message = "Refresh Successful."
		response.Payload = tokens
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.JSON(http.StatusOK, response)
	}
}

func (auth *AuthInterface) ResetPassword(ctx *gin.Context) {
	// To be done later.
	response := value_objects.Response{}

	response.Status = true
	response.Message = "Reset Password Hit."
	response.Payload = nil

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, response)
}
