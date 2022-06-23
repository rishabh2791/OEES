package persistance

import (
	"oees/domain/entity"
	"oees/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type shiftRepo struct {
	db     *gorm.DB
	logger hclog.Logger
}

var _ repository.ShiftRepository = &shiftRepo{}

func newshiftRepo(db *gorm.DB, logger hclog.Logger) *shiftRepo {
	return &shiftRepo{
		db:     db,
		logger: logger,
	}
}

func (shiftRepo *shiftRepo) Create(shift *entity.Shift) (*entity.Shift, error) {
	validationErr := shift.Validate("create")
	if validationErr != nil {
		return nil, validationErr
	}
	creationErr := shiftRepo.db.Create(&shift).Error
	return shift, creationErr
}

func (shiftRepo *shiftRepo) Get(id string) (*entity.Shift, error) {
	shift := entity.Shift{}
	getErr := shiftRepo.db.
		Preload("Plant.CreatedBy").
		Preload("Plant.CreatedBy.UserRole").
		Preload("Plant.UpdatedBy").
		Preload("Plant.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).Take(&shift).Error
	return &shift, getErr
}

func (shiftRepo *shiftRepo) List(conditions string) ([]entity.Shift, error) {
	shifts := []entity.Shift{}
	getErr := shiftRepo.db.
		Preload("Plant.CreatedBy").
		Preload("Plant.CreatedBy.UserRole").
		Preload("Plant.UpdatedBy").
		Preload("Plant.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&shifts).Error
	return shifts, getErr
}

func (shiftRepo *shiftRepo) Update(id string, update *entity.Shift) (*entity.Shift, error) {
	existingShift := entity.Shift{}
	getErr := shiftRepo.db.Where("id = ?", id).Take(&existingShift).Error
	if getErr != nil {
		return nil, getErr
	}
	updationErr := shiftRepo.db.Table(update.Tablename()).Where("id = ?", id).Updates(&update).Error
	if updationErr != nil {
		return nil, updationErr
	}
	updated := entity.Shift{}
	shiftRepo.db.
		Preload("Plant.CreatedBy").
		Preload("Plant.CreatedBy.UserRole").
		Preload("Plant.UpdatedBy").
		Preload("Plant.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).Take(&updated)
	return &updated, nil
}
