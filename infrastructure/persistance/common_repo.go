package persistance

import (
	"oees/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type commonRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.CommonRepository = &commonRepo{}

func newCommonRepo(db *gorm.DB, logger hclog.Logger) *commonRepo {
	return &commonRepo{
		DB:     db,
		Logger: logger,
	}
}

func (commonRepo *commonRepo) GetTables() ([]string, error) {
	tables := []string{}

	getErr := commonRepo.DB.Table("information_schema.tables").Where("table_schema = ?", "oees").Pluck("table_name", &tables).Error

	if getErr != nil {
		return nil, getErr
	}

	return tables, nil
}
