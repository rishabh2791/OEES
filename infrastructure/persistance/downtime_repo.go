package persistance

import (
	"errors"
	"fmt"
	"oees/domain/entity"
	"oees/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type downtimeRepo struct {
	db     *gorm.DB
	logger hclog.Logger
}

var _ repository.DowntimeRepository = &downtimeRepo{}

func newdowntimeRepo(db *gorm.DB, logger hclog.Logger) *downtimeRepo {
	return &downtimeRepo{
		db:     db,
		logger: logger,
	}
}

func (downtimeRepo *downtimeRepo) Create(downtime *entity.Downtime) (*entity.Downtime, error) {
	validationErr := downtime.Validate("create")
	if validationErr != nil {
		return nil, validationErr
	}
	checkExistingDowntime := []entity.Downtime{}
	existingQuery := fmt.Sprintf("SELECT * FROM `downtimes` WHERE start_time < '%s' AND (end_time > '%s' OR end_time IS NULL);", downtime.StartTime, downtime.StartTime)
	downtimeRepo.db.Raw(existingQuery).Find(&checkExistingDowntime)
	if len(checkExistingDowntime) == 0 {
		creationErr := downtimeRepo.db.Create(&downtime).Error
		return downtime, creationErr
	}
	return nil, errors.New("Existing Downtime")
}

func (downtimeRepo *downtimeRepo) Get(id string) (*entity.Downtime, error) {
	downtime := entity.Downtime{}
	getErr := downtimeRepo.db.
		Preload("Line.CreatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy").
		Preload("Line.UpdatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).
		Where("id = ?", id).Take(&downtime).Error
	return &downtime, getErr
}

func (downtimeRepo *downtimeRepo) List(conditions string) ([]entity.Downtime, error) {
	downtimes := []entity.Downtime{}
	getErr := downtimeRepo.db.
		Preload("Line.CreatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy").
		Preload("Line.UpdatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).
		Where(conditions).Find(&downtimes).Error
	return downtimes, getErr
}

func (downtimeRepo *downtimeRepo) Update(id string, update *entity.Downtime) (*entity.Downtime, error) {
	existingDowntime := entity.Downtime{}
	getErr := downtimeRepo.db.Where("id = ?", id).Take(&existingDowntime).Error
	if getErr != nil {
		return nil, getErr
	}
	updationErr := downtimeRepo.db.Table(update.Tablename()).Where("id = ?", id).Updates(&update).Error
	if updationErr != nil {
		return nil, updationErr
	}
	updated := entity.Downtime{}
	downtimeRepo.db.
		Preload("Line.CreatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy").
		Preload("Line.UpdatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).
		Where("id = ?", id).Take(&updated)
	return &updated, nil
}
