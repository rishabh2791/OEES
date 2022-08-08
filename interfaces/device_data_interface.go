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

type DeviceDataInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
	// channel  chan *entity.DeviceData
}

func NewDeviceDataInterface(appStore *application.AppStore, logger hclog.Logger) *DeviceDataInterface {
	// channel := make(chan *entity.DeviceData, 10)
	return &DeviceDataInterface{
		appStore: appStore,
		logger:   logger,
		// channel:  channel,
	}
}

// var wsupgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// func (deviceDataInterface *DeviceDataInterface) wshandler(w http.ResponseWriter, r *http.Request, weight chan<- *entity.DeviceData) {
// 	connections := map[string]*websocket.Conn{}
// 	conn, err := wsupgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		return
// 	}
// 	connections[conn.RemoteAddr().String()] = conn
// 	log.Println(connections)

// 	newWeight := <-deviceDataInterface.channel
// 	log.Println(newWeight)

// 	for {
// 		t, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			break
// 		}
// 		conn.WriteMessage(t, msg)
// 	}
// }

func (deviceDataInterface *DeviceDataInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	// Get new entry details from request body.
	model := entity.DeviceData{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create entry in database.
	created, creationErr := deviceDataInterface.appStore.DeviceDataApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	// deviceDataInterface.channel <- created
	response.Status = true
	response.Message = "Device Data Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (deviceDataInterface *DeviceDataInterface) List(ctx *gin.Context) {
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

	data, err := deviceDataInterface.appStore.DeviceDataApp.List(utilities.ConvertJSONToSQL(conditions))
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Device Data Found"
	response.Payload = data

	ctx.JSON(http.StatusOK, response)
}

func (deviceDataInterface *DeviceDataInterface) TotalDeviceData(ctx *gin.Context) {
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

	data, err := deviceDataInterface.appStore.DeviceDataApp.TotalDeviceData(utilities.ConvertJSONToSQL(conditions))
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Device Data Found"
	response.Payload = data

	ctx.JSON(http.StatusOK, response)
}

// func (deviceDataInterface *DeviceDataInterface) WebSocketHandler(ctx *gin.Context) {
// 	deviceDataInterface.wshandler(ctx.Writer, ctx.Request, deviceDataInterface.channel)
// }
