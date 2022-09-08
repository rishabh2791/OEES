package persistance

import (
	"oees/domain/entity"
	"oees/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type presetDowntimeRepo struct {
	db     *gorm.DB
	logger hclog.Logger
}

var _ repository.PresetDowntimeRepository = &presetDowntimeRepo{}

func newPresetDowntimeRepo(db *gorm.DB, logger hclog.Logger) *presetDowntimeRepo {
	return &presetDowntimeRepo{
		db:     db,
		logger: logger,
	}
}

func (presetDowntimeRepo *presetDowntimeRepo) Create(downtime *entity.PresetDowntime) (*entity.PresetDowntime, error) {
	validationErr := downtime.Validate("create")
	if validationErr != nil {
		return nil, validationErr
	}
	creationErr := presetDowntimeRepo.db.Create(&downtime).Error
	return downtime, creationErr
}
func (presetDowntimeRepo *presetDowntimeRepo) List(conditions string) ([]entity.PresetDowntime, error) {
	presetDowntimes := []entity.PresetDowntime{}
	getErr := presetDowntimeRepo.db.Where(conditions).Find(&presetDowntimes).Error
	return presetDowntimes, getErr
}
