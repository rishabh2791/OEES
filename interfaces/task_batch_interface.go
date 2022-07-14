package interfaces

import (
	"encoding/json"
	"net/http"
	"oees/application"
	"oees/domain/entity"
	"oees/domain/value_objects"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

type TaskBatchInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func newTaskBatchInterface(appStore *application.AppStore, logger hclog.Logger) *TaskBatchInterface {
	return &TaskBatchInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (taskBatchInterface *TaskBatchInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	// Get new entry details from request body.
	model := entity.TaskBatch{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create entry in database.
	created, creationErr := taskBatchInterface.appStore.TaskBatchApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "Task Batch Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (taskBatchInterface *TaskBatchInterface) List(ctx *gin.Context) {
	response := value_objects.Response{}

	id := ctx.Param("taskID")
	task, err := taskBatchInterface.appStore.TaskBatchApp.List(id)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Task Batches Found"
	response.Payload = task

	ctx.JSON(http.StatusOK, response)
}

func (taskBatchInterface *TaskBatchInterface) Update(ctx *gin.Context) {
	response := value_objects.Response{}
	id := ctx.Param("taskID")

	// Get new entry details from request body.
	model := entity.TaskBatch{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create entry in database.
	updated, creationErr := taskBatchInterface.appStore.TaskBatchApp.Update(id, &model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "Task Batch Updated."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}
