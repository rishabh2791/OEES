package persistance

import (
	"oees/domain/entity"
	"oees/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type deviceDataRepo struct {
	db     *gorm.DB
	logger hclog.Logger
}

var _ repository.DeviceDataRepository = &deviceDataRepo{}

func newDeviceDataRepo(db *gorm.DB, logger hclog.Logger) *deviceDataRepo {
	return &deviceDataRepo{
		db:     db,
		logger: logger,
	}
}

func (deviceDataRepo *deviceDataRepo) Create(deviceData *entity.DeviceData) (*entity.DeviceData, error) {
	validationErr := deviceData.Validate("Create")
	if validationErr != nil {
		return nil, validationErr
	}
	creationErr := deviceDataRepo.db.Create(&deviceData).Error
	return deviceData, creationErr
}

func (deviceDataRepo *deviceDataRepo) List(conditions string) ([]entity.DeviceData, error) {
	deviceData := []entity.DeviceData{}
	getErr := deviceDataRepo.db.
		Preload("Device.CreatedBy").
		Preload("Device.CreatedBy.UserRole").
		Preload("Device.UpdatedBy").
		Preload("Device.UpdatedBy.UserRole").
		Preload("Device.Line.CreatedBy").
		Preload("Device.Line.CreatedBy.UserRole").
		Preload("Device.Line.UpdatedBy").
		Preload("Device.Line.UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&deviceData).Error
	return deviceData, getErr
}
