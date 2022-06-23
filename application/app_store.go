package application

import "oees/infrastructure/persistance"

type AppStore struct {
	repoStore         *persistance.RepoStore
	AuthApp           *authApp
	CommonApp         *commonApp
	DeviceApp         *deviceApp
	DeviceDataApp     *deviceDataApp
	DowntimeApp       *downtimeApp
	DowntimePresetApp *presetDowntimeApp
	JobApp            *jobApp
	LineApp           *lineApp
	PlantApp          *plantApp
	ShiftApp          *shiftApp
	SKUApp            *skuApp
	SKUSpeedApp       *skuSpeedApp
	TaskApp           *taskApp
	UserRoleApp       *userRoleApp
	UserApp           *userApp
	UserRoleAccessApp *userRoleAccessApp
}

func NewAppStore(repoStore *persistance.RepoStore) *AppStore {
	return &AppStore{
		repoStore:         repoStore,
		AuthApp:           newAuthApp(repoStore.AuthRepo),
		CommonApp:         newCommonApp(repoStore.CommonRepo),
		DeviceApp:         newdeviceApp(repoStore.DeviceRepo),
		DeviceDataApp:     newdeviceDataApp(repoStore.DeviceDataRepo),
		DowntimeApp:       newdowntimeApp(repoStore.DowntimeRepo),
		DowntimePresetApp: newPresetDowntimeApp(repoStore.DowntimePresetRepo),
		JobApp:            newJobApp(repoStore.JobRepo),
		LineApp:           newlineApp(repoStore.LineRepo),
		PlantApp:          newplantApp(repoStore.PlantRepo),
		ShiftApp:          newshiftApp(repoStore.ShiftRepo),
		SKUApp:            newskuApp(repoStore.SkuRepo),
		SKUSpeedApp:       newskuSpeedApp(repoStore.SkuSpeedRepo),
		TaskApp:           newTaskApp(repoStore.TaskRepo),
		UserRoleApp:       newuserRoleApp(repoStore.UserRoleRepo),
		UserApp:           newUserApp(repoStore.UserRepo),
		UserRoleAccessApp: NewUserRoleAccessApp(repoStore.UserRoleAccessRepo),
	}
}
