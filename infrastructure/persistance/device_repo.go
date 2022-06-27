package persistance

import (
	"oees/domain/entity"
	"oees/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type deviceRepo struct {
	db     *gorm.DB
	logger hclog.Logger
}

var _ repository.DeviceRepository = &deviceRepo{}

func newDeviceRepo(db *gorm.DB, logger hclog.Logger) *deviceRepo {
	return &deviceRepo{
		db:     db,
		logger: logger,
	}
}

func (deviceRepo *deviceRepo) Create(device *entity.Device) (*entity.Device, error) {
	validationErr := device.Validate("create")
	if validationErr != nil {
		return nil, validationErr
	}
	creationErr := deviceRepo.db.Create(&device).Error
	return device, creationErr
}

func (deviceRepo *deviceRepo) Get(id string) (*entity.Device, error) {
	device := entity.Device{}
	getErr := deviceRepo.db.
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload("Line.CreatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy").
		Preload("Line.UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).Take(&device).Error
	return &device, getErr
}

func (deviceRepo *deviceRepo) List(conditions string) ([]entity.Device, error) {
	devices := []entity.Device{}
	getErr := deviceRepo.db.
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload("Line.CreatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy").
		Preload("Line.UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&devices).Error
	return devices, getErr
}

func (deviceRepo *deviceRepo) Update(id string, update *entity.Device) (*entity.Device, error) {
	existingDevice := entity.Device{}
	getErr := deviceRepo.db.Where("id = ?", id).Take(&existingDevice).Error
	if getErr != nil {
		return nil, getErr
	}
	updationErr := deviceRepo.db.Table(update.Tablename()).Where("id = ?", id).Updates(&update).Error
	if updationErr != nil {
		return nil, updationErr
	}
	updated := entity.Device{}
	deviceRepo.db.
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload("Line.CreatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy").
		Preload("Line.UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).Take(&updated)
	return &updated, nil
}
