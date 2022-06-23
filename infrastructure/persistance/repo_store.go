package persistance

import (
	"net/url"
	"oees/domain/entity"
	"oees/infrastructure/utilities"
	"os"

	"github.com/go-redis/redis"
	"github.com/hashicorp/go-hclog"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type RepoStore struct {
	DB                 *gorm.DB
	warehouseDB        *gorm.DB
	Cache              *redis.Client
	Logger             hclog.Logger
	AuthRepo           *authRepo
	CommonRepo         *commonRepo
	DeviceRepo         *deviceRepo
	DeviceDataRepo     *deviceDataRepo
	DowntimeRepo       *downtimeRepo
	DowntimePresetRepo *presetDowntimeRepo
	JobRepo            *jobRepo
	LineRepo           *lineRepo
	PlantRepo          *plantRepo
	ShiftRepo          *shiftRepo
	SkuRepo            *skuRepo
	SkuSpeedRepo       *skuSpeedRepo
	TaskRepo           *taskRepo
	UserRepo           *userRepo
	UserRoleRepo       *userRoleRepo
	UserRoleAccessRepo *UserRoleAccessRepo
}

func NewRepoStore(config *utilities.ServerConfig, logging hclog.Logger) (*RepoStore, error) {
	repoStore := RepoStore{}
	databaseConfig := utilities.NewDatabaseConfig()

	cacheStore, cacheError := NewCacheStore(*config)
	if cacheError != nil {
		logging.Error(cacheError.Error())
		os.Exit(1)
	}

	mysqlURL := databaseConfig.DbUser + ":" + databaseConfig.DbPassword + "@tcp(" + databaseConfig.DbHost + ":" + databaseConfig.DbPort + ")/" + databaseConfig.DbName + "?parseTime=True"
	gormDB, gormErr := gorm.Open(mysql.Open(mysqlURL), &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Silent),
		QueryFields:          true,
		FullSaveAssociations: true,
	})
	if gormErr != nil {
		return nil, gormErr
	}
	sqlDB, _ := gormDB.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(10000)

	// MSSQL Connection
	username := url.QueryEscape(databaseConfig.WarehouseUser)
	password := url.QueryEscape(databaseConfig.WarehousePassword)
	sqlURL := "sqlserver://" + username + ":" + password + "@" + databaseConfig.WarehouseHost + ":1433?database=" + databaseConfig.WarehouseDBName
	remoteSQLDB, _ := gorm.Open(sqlserver.Open(sqlURL), &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Silent),
		QueryFields:          true,
		FullSaveAssociations: true,
	})
	// if sqlErr != nil {
	// 	return nil, sqlErr
	// }
	repoStore.warehouseDB = remoteSQLDB

	repoStore.DB = gormDB
	repoStore.Logger = logging
	repoStore.Cache = cacheStore.RedisClient
	repoStore.AuthRepo = newAuthRepo(logging, cacheStore.RedisClient, config)
	repoStore.CommonRepo = newCommonRepo(gormDB, logging)
	repoStore.DeviceRepo = newDeviceRepo(gormDB, logging)
	repoStore.DeviceDataRepo = newDeviceDataRepo(gormDB, logging)
	repoStore.DowntimeRepo = newdowntimeRepo(gormDB, logging)
	repoStore.DowntimePresetRepo = newPresetDowntimeRepo(gormDB, logging)
	repoStore.JobRepo = newJobRepo(gormDB, logging)
	repoStore.LineRepo = newlineRepo(gormDB, logging)
	repoStore.PlantRepo = newplantRepo(gormDB, logging)
	repoStore.ShiftRepo = newshiftRepo(gormDB, logging)
	repoStore.SkuRepo = newskuRepo(gormDB, logging)
	repoStore.TaskRepo = NewTaskRepo(gormDB, logging)
	repoStore.SkuSpeedRepo = newskuSpeedRepo(gormDB, logging)
	repoStore.UserRepo = newuserRepo(gormDB, logging)
	repoStore.UserRoleRepo = newUserRoleRepo(gormDB, logging)
	repoStore.UserRoleAccessRepo = NewUserRoleAccessRepo(gormDB, logging)

	return &repoStore, nil
}

func (repoStore *RepoStore) Migrate() error {
	return repoStore.DB.AutoMigrate(
		&entity.UserRole{},
		&entity.User{},
		&entity.UserRoleAccess{},
		&entity.Plant{},
		&entity.Line{},
		&entity.Device{},
		&entity.DeviceData{},
		&entity.Shift{},
		&entity.Downtime{},
		&entity.SKU{},
		&entity.SKUSpeed{},
		&entity.Task{},
		&entity.Job{},
		&entity.PresetDowntime{},
	)
}
