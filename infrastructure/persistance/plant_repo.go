package persistance

import (
	"oees/domain/entity"
	"oees/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type plantRepo struct {
	db     *gorm.DB
	logger hclog.Logger
}

var _ repository.PlantRepository = &plantRepo{}

func newplantRepo(db *gorm.DB, logger hclog.Logger) *plantRepo {
	return &plantRepo{
		db:     db,
		logger: logger,
	}
}

func (plantRepo *plantRepo) Create(plant *entity.Plant) (*entity.Plant, error) {
	validationErr := plant.Validate("create")
	if validationErr != nil {
		return nil, validationErr
	}
	creationErr := plantRepo.db.Create(&plant).Error
	return plant, creationErr
}

func (plantRepo *plantRepo) Get(id string) (*entity.Plant, error) {
	plant := entity.Plant{}
	getErr := plantRepo.db.
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).
		Where("id = ?", id).Take(&plant).Error
	return &plant, getErr
}

func (plantRepo *plantRepo) List(conditions string) ([]entity.Plant, error) {
	plants := []entity.Plant{}
	getErr := plantRepo.db.
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).
		Where(conditions).Find(&plants).Error
	return plants, getErr
}

func (plantRepo *plantRepo) Update(id string, update *entity.Plant) (*entity.Plant, error) {
	existingPlant := entity.Plant{}
	getErr := plantRepo.db.Where("id = ?", id).Take(&existingPlant).Error
	if getErr != nil {
		return nil, getErr
	}
	updationErr := plantRepo.db.Table(update.Tablename()).Where("id = ?", id).Updates(&update).Error
	if updationErr != nil {
		return nil, updationErr
	}
	updated := entity.Plant{}
	plantRepo.db.
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).
		Where("id = ?", id).Take(&updated)
	return &updated, getErr
}
