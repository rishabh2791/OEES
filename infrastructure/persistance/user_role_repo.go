package persistance

import (
	"oees/domain/entity"
	"oees/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRoleRepo struct {
	db     *gorm.DB
	logger hclog.Logger
}

var _ repository.UserRoleRepository = &userRoleRepo{}

func newUserRoleRepo(db *gorm.DB, logger hclog.Logger) *userRoleRepo {
	return &userRoleRepo{
		db:     db,
		logger: logger,
	}
}

func (userRoleRepo *userRoleRepo) Create(userRole *entity.UserRole) (*entity.UserRole, error) {
	validationErr := userRole.Validate("create")
	if validationErr != nil {
		return nil, validationErr
	}
	creationErr := userRoleRepo.db.Create(&userRole).Error
	return userRole, creationErr
}

func (userRoleRepo *userRoleRepo) List(conditions string) ([]entity.UserRole, error) {
	userRoles := []entity.UserRole{}
	getErr := userRoleRepo.db.
		Preload(clause.Associations).Where(conditions).Find(&userRoles).Error
	return userRoles, getErr
}
