package persistance

import (
	"oees/domain/entity"
	"oees/domain/repository"
	"oees/domain/value_objects"

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
	// If there are counts, from posted JSON, get deviceID and use it to find device details and hence the Line ID.
	// Check if there are any open downtimes for this lineID and if yes, close it.
	if deviceData.Value > 0 {
		deviceID := deviceData.DeviceID
		device := entity.Device{}
		getErr := deviceDataRepo.db.
			Preload(clause.Associations).Where("id = ?", deviceID).Take(&device).Error
		if getErr != nil {
			return nil, getErr
		}

		lineID := device.LineID
		downtimes := []entity.Downtime{}
		getErr = deviceDataRepo.db.
			Preload(clause.Associations).
			Where("line_id LIKE ? AND end_time IS NULL", lineID).Find(&downtimes).Error

		if len(downtimes) != 0 {
			for _, downtime := range downtimes {
				query := "UPDATE downtimes SET end_time = UTC_TIMESTAMP() WHERE id LIKE \"" + downtime.ID + "\""
				updateErr := deviceDataRepo.db.Exec(query).Error
				if updateErr != nil {
					return nil, updateErr
				}
			}
		}
	}

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
