package persistance

import (
	"oees/domain/entity"
	"oees/domain/repository"
	"oees/domain/value_objects"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
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
	getErr := deviceDataRepo.db.Where(conditions).Find(&deviceData).Error
	return deviceData, getErr
}

func (deviceDataRepo *deviceDataRepo) TotalDeviceData(conditions string) (*value_objects.TotalDeviceData, error) {
	deviceData := value_objects.TotalDeviceData{}

	sqlQuery := "SELECT sum(value) as value, device_id FROM device_data WHERE " + conditions
	rows, getErr := deviceDataRepo.db.Raw(sqlQuery).Rows()
	if getErr != nil {
		return nil, getErr
	}
	defer rows.Close()
	for rows.Next() {
		deviceDataRepo.db.ScanRows(rows, &deviceData)
	}
	return &deviceData, nil
}
