package persistance

import (
	"oees/domain/entity"
	"oees/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
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

func (userRoleRepo *userRoleRepo) Get(userRoleID string) (*entity.UserRole, error) {
	userRole := entity.UserRole{}

	getErr := userRoleRepo.db.Where("id = ?", userRoleID).Take(&userRole).Error

	return &userRole, getErr
}

func (userRoleRepo *userRoleRepo) List(conditions string) ([]entity.UserRole, error) {
	userRoles := []entity.UserRole{}
	getErr := userRoleRepo.db.Where(conditions).Find(&userRoles).Error
	return userRoles, getErr
}
