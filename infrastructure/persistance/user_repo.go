package persistance

import (
	"oees/domain/entity"
	"oees/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepo struct {
	db     *gorm.DB
	logger hclog.Logger
}

var _ repository.UserRepository = &userRepo{}

func newuserRepo(db *gorm.DB, logger hclog.Logger) *userRepo {
	return &userRepo{
		db:     db,
		logger: logger,
	}
}

func (userRepo *userRepo) Create(user *entity.User) (*entity.User, error) {
	validationErr := user.Validate("create")
	if validationErr != nil {
		return nil, validationErr
	}
	creationErr := userRepo.db.Create(&user).Error
	return user, creationErr
}

func (userRepo *userRepo) Get(username string) (*entity.User, error) {
	user := entity.User{}
	getErr := userRepo.db.
		Preload(clause.Associations).Where("username = ?", username).Take(&user).Error
	return &user, getErr
}

func (userRepo *userRepo) List(conditions string) ([]entity.User, error) {
	users := []entity.User{}
	getErr := userRepo.db.Preload(clause.Associations).Where(conditions).Find(&users).Error
	return users, getErr
}

func (userRepo *userRepo) Update(username string, update *entity.User) (*entity.User, error) {
	existingUser := entity.User{}

	err := userRepo.db.Where("username = ?", username).Take(&existingUser).Error
	if err != nil {
		return nil, err
	}

	updationErr := userRepo.db.Table(existingUser.Tablename()).Where("username = ?", username).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.User{}
	userRepo.db.Where("username = ?", username).Take(&updated)

	return &updated, nil
}
