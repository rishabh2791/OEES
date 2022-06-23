package persistance

import (
	"oees/domain/entity"
	"oees/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type skuSpeedRepo struct {
	db     *gorm.DB
	logger hclog.Logger
}

var _ repository.SKUSpeedRepository = &skuSpeedRepo{}

func newskuSpeedRepo(db *gorm.DB, logger hclog.Logger) *skuSpeedRepo {
	return &skuSpeedRepo{
		db:     db,
		logger: logger,
	}
}

func (skuSpeedRepo *skuSpeedRepo) Create(skuSpeed *entity.SKUSpeed) (*entity.SKUSpeed, error) {
	validationErr := skuSpeed.Validate("create")
	if validationErr != nil {
		return nil, validationErr
	}
	creationErr := skuSpeedRepo.db.Create(&skuSpeed).Error
	return skuSpeed, creationErr
}

func (skuSpeedRepo *skuSpeedRepo) Get(id string) (*entity.SKUSpeed, error) {
	skuSpeed := entity.SKUSpeed{}
	getErr := skuSpeedRepo.db.
		Preload("Line.Plant").
		Preload("Line.Plant.CreatedBy").
		Preload("Line.Plant.CreatedBy.UserRole").
		Preload("Line.Plant.UpdatedBy").
		Preload("Line.Plant.UpdatedBy.UserRole").
		Preload("Line.CreatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy").
		Preload("Line.UpdatedBy.UserRole").
		Preload("SKU.Plant").
		Preload("SKU.Plant.CreatedBy").
		Preload("SKU.Plant.CreatedBy.UserRole").
		Preload("SKU.Plant.UpdatedBy").
		Preload("SKU.Plant.UpdatedBy.UserRole").
		Preload("SKU.CreatedBy").
		Preload("SKU.CreatedBy.UserRole").
		Preload("SKU.UpdatedBy").
		Preload("SKU.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).Take(&skuSpeed).Error
	return &skuSpeed, getErr
}

func (skuSpeedRepo *skuSpeedRepo) List(conditions string) ([]entity.SKUSpeed, error) {
	skuSpeeds := []entity.SKUSpeed{}
	getErr := skuSpeedRepo.db.
		Preload("Line.Plant").
		Preload("Line.Plant.CreatedBy").
		Preload("Line.Plant.CreatedBy.UserRole").
		Preload("Line.Plant.UpdatedBy").
		Preload("Line.Plant.UpdatedBy.UserRole").
		Preload("Line.CreatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy").
		Preload("Line.UpdatedBy.UserRole").
		Preload("SKU.Plant").
		Preload("SKU.Plant.CreatedBy").
		Preload("SKU.Plant.CreatedBy.UserRole").
		Preload("SKU.Plant.UpdatedBy").
		Preload("SKU.Plant.UpdatedBy.UserRole").
		Preload("SKU.CreatedBy").
		Preload("SKU.CreatedBy.UserRole").
		Preload("SKU.UpdatedBy").
		Preload("SKU.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&skuSpeeds).Error
	return skuSpeeds, getErr
}

func (skuSpeedRepo *skuSpeedRepo) Update(id string, update *entity.SKUSpeed) (*entity.SKUSpeed, error) {
	existingSKUSpeed := entity.SKUSpeed{}
	getErr := skuSpeedRepo.db.Preload(clause.Associations).Where("id = ?", id).Take(&existingSKUSpeed).Error
	if getErr != nil {
		return nil, getErr
	}
	updationErr := skuSpeedRepo.db.Table(update.Tablename()).Where("id = ?", id).Updates(&update).Error
	if updationErr != nil {
		return nil, updationErr
	}
	updated := entity.SKUSpeed{}
	skuSpeedRepo.db.
		Preload("Line.Plant").
		Preload("Line.Plant.CreatedBy").
		Preload("Line.Plant.CreatedBy.UserRole").
		Preload("Line.Plant.UpdatedBy").
		Preload("Line.Plant.UpdatedBy.UserRole").
		Preload("Line.CreatedBy").
		Preload("Line.CreatedBy.UserRole").
		Preload("Line.UpdatedBy").
		Preload("Line.UpdatedBy.UserRole").
		Preload("SKU.Plant").
		Preload("SKU.Plant.CreatedBy").
		Preload("SKU.Plant.CreatedBy.UserRole").
		Preload("SKU.Plant.UpdatedBy").
		Preload("SKU.Plant.UpdatedBy.UserRole").
		Preload("SKU.CreatedBy").
		Preload("SKU.CreatedBy.UserRole").
		Preload("SKU.UpdatedBy").
		Preload("SKU.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).Take(&updated)
	return update, nil
}
