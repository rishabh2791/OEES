package interfaces

import (
	"oees/application"

	"github.com/hashicorp/go-hclog"
)

type InterfaceStore struct {
	appStore                *application.AppStore
	logger                  hclog.Logger
	AuthInterface           *AuthInterface
	CommonInterface         *CommonInterface
	DeviceInterface         *DeviceInterface
	DeviceDataInterface     *DeviceDataInterface
	DowntimeInterface       *DowntimeInterface
	DowntimePresetInterface *DowntimePresetInterface
	JobInterface            *JobInterface
	LineInterface           *LineInterface
	PlantInterface          *PlantInterface
	ShiftInterface          *ShiftInterface
	SKUInterface            *SKUInterface
	SKUSpeedInterface       *SKUSpeedInterface
	TaskInterface           *TaskInterface
	UserRoleInterface       *UserRoleInterface
	UserInterface           *UserInterface
	UserRoleAccessInterface *UserRoleAccessInterface
}

func NewInterfaceStore(appStore *application.AppStore, logger hclog.Logger) *InterfaceStore {
	return &InterfaceStore{
		appStore:                appStore,
		logger:                  logger,
		AuthInterface:           NewAuthInterface(appStore, logger),
		CommonInterface:         NewCommonInterface(appStore, logger),
		DeviceInterface:         NewDeviceInterface(appStore, logger),
		DeviceDataInterface:     NewDeviceDataInterface(appStore, logger),
		DowntimeInterface:       NewDowntimeInterface(appStore, logger),
		DowntimePresetInterface: NewDowntimePresetInterface(appStore, logger),
		JobInterface:            (*JobInterface)(NewJobInterface(appStore, logger)),
		LineInterface:           NewlineInterface(appStore, logger),
		PlantInterface:          NewplantInterface(appStore, logger),
		ShiftInterface:          NewshiftInterface(appStore, logger),
		SKUInterface:            NewskuInterface(appStore, logger),
		SKUSpeedInterface:       NewskuSpeedInterface(appStore, logger),
		TaskInterface:           NewTaskInterface(appStore, logger),
		UserRoleInterface:       NewuserRoleInterface(appStore, logger),
		UserInterface:           NewUserInterface(appStore, logger),
		UserRoleAccessInterface: NewUserRoleAccessInterface(appStore, logger),
	}
}
