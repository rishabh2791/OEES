package persistance

import (
	"oees/domain/entity"
	"oees/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type lineRepo struct {
	db     *gorm.DB
	logger hclog.Logger
}

var _ repository.LineRepository = &lineRepo{}

func newlineRepo(db *gorm.DB, logger hclog.Logger) *lineRepo {
	return &lineRepo{
		db:     db,
		logger: logger,
	}
}

func (lineRepo *lineRepo) Create(line *entity.Line) (*entity.Line, error) {
	validationErr := line.Validate("create")
	if validationErr != nil {
		return nil, validationErr
	}
	creationErr := lineRepo.db.Create(&line).Error
	return line, creationErr
}

func (lineRepo *lineRepo) Get(id string) (*entity.Line, error) {
	line := entity.Line{}
	getErr := lineRepo.db.Where("id = ?", id).Take(&line).Error
	return &line, getErr
}

func (lineRepo *lineRepo) List(conditions string) ([]entity.Line, error) {
	lines := []entity.Line{}
	getErr := lineRepo.db.Where(conditions).Find(&lines).Error
	return lines, getErr
}

func (lineRepo *lineRepo) Update(id string, update *entity.Line) (*entity.Line, error) {
	existingLine := entity.Line{}
	getErr := lineRepo.db.Where("id = ?", id).Take(&existingLine).Error
	if getErr != nil {
		return nil, getErr
	}
	updationErr := lineRepo.db.Table(update.Tablename()).Where("id = ?", id).Updates(&update).Error
	if updationErr != nil {
		return nil, updationErr
	}
	updated := entity.Line{}
	lineRepo.db.Where("id = ?", id).Take(&updated)
	return &updated, nil
}
