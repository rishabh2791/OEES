package persistance

import (
	"oees/domain/entity"
	"oees/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type skuRepo struct {
	db     *gorm.DB
	logger hclog.Logger
}

var _ repository.SKURepository = &skuRepo{}

func newskuRepo(db *gorm.DB, logger hclog.Logger) *skuRepo {
	return &skuRepo{
		db:     db,
		logger: logger,
	}
}

func (skuRepo *skuRepo) Create(sku *entity.SKU) (*entity.SKU, error) {
	validationErr := sku.Validate("create")
	if validationErr != nil {
		return nil, validationErr
	}
	creationErr := skuRepo.db.Create(&sku).Error
	return sku, creationErr
}

func (skuRepo *skuRepo) Get(id string) (*entity.SKU, error) {
	sku := entity.SKU{}
	getErr := skuRepo.db.
		Preload("Plant.CreatedBy").
		Preload("Plant.CreatedBy.UserRole").
		Preload("Plant.UpdatedBy").
		Preload("Plant.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).Take(&sku).Error
	return &sku, getErr
}

func (skuRepo *skuRepo) List(conditions string) ([]entity.SKU, error) {
	skus := []entity.SKU{}
	getErr := skuRepo.db.
		Preload("Plant.CreatedBy").
		Preload("Plant.CreatedBy.UserRole").
		Preload("Plant.UpdatedBy").
		Preload("Plant.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&skus).Error
	return skus, getErr
}

func (skuRepo *skuRepo) Update(id string, update *entity.SKU) (*entity.SKU, error) {
	existingSKU := entity.SKU{}

	err := skuRepo.db.Where("id = ?", id).Take(&existingSKU).Error
	if err != nil {
		return nil, err
	}

	updationErr := skuRepo.db.Table(existingSKU.Tablename()).Where("id = ?", id).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.SKU{}
	skuRepo.db.
		Preload("Plant.CreatedBy").
		Preload("Plant.CreatedBy.UserRole").
		Preload("Plant.UpdatedBy").
		Preload("Plant.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).
		Where("id = ?", id).Take(&updated)

	return &updated, nil
}
