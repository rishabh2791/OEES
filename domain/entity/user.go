package entity

import (
	"oees/domain/value_objects"
	"oees/infrastructure/utilities"
	"oees/infrastructure/utilities/security"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	value_objects.BaseModel
	UserRoleID string    `json:"user_role_id" gorm:"size:191;not null;"`
	UserRole   *UserRole `json:"user_role"`
	Username   string    `json:"username" gorm:"size:20;not null;unique;primaryKey;"`
	FirstName  string    `json:"first_name" gorm:"size:100;not null;"`
	LastName   string    `json:"last_name" gorm:"size:100;"`
	Password   string    `json:"password" gorm:"size:200;not null;"`
	Email      string    `json:"email" gorm:"size:100;not null;"`
	ProfilePic string    `json:"profile_pic" gorm:"size:1000;not null;default:'public/profile_pics/default.jpg'"`
	Active     bool      `json:"active" gorm:"default:true;"`
	SecretKey  string    `json:"secret_key" gorm:"size:191;"`
}

var _ Model = &User{}

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.SecretKey = strings.ReplaceAll(uuid.NewString(), "-", "")
	if user.ProfilePic == "" || len(user.ProfilePic) == 0 {
		user.ProfilePic = "public/profile_pics/default.jpg"
	}
	hashedPassword, passError := security.Hash(user.Password)
	if passError != nil {
		return passError
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) Validate(action string) error {
	errors := map[string]interface{}{}
	switch action {
	case "login":
		if user.Username == "" || len(user.Username) == 0 {
			errors["username"] = "Username Missing"
		}
		if user.Password == "" || len(user.Password) == 0 {
			errors["password"] = "Password Missing"
		}
	case "superuser":
		if user.Username == "" || len(user.Username) == 0 {
			errors["username"] = "Username Missing"
		}
		if user.FirstName == "" || len(user.FirstName) == 0 {
			errors["first_name"] = "First Name Missing"
		}
		if user.Email == "" || len(user.Email) == 0 {
			errors["email"] = "EMail Missing"
		}
		if user.Password == "" || len(user.Password) == 0 {
			errors["password"] = "Password Missing"
		}
	case "reset_password":
		if user.Username == "" || len(user.Username) == 0 {
			errors["username"] = "Username Missing"
		}
		if user.Password == "" || len(user.Password) == 0 {
			errors["password"] = "Password Missing"
		}
	default:
		if user.Username == "" || len(user.Username) == 0 {
			errors["username"] = "Username Missing"
		}
		if user.FirstName == "" || len(user.FirstName) == 0 {
			errors["first_name"] = "First Name Missing"
		}
		if user.Email == "" || len(user.Email) == 0 {
			errors["email"] = "EMail Missing"
		}
		if user.Password == "" || len(user.Password) == 0 {
			errors["password"] = "Password Missing"
		}
	}
	return utilities.ConvertMapToError(errors)
}

func (user *User) Tablename() string {
	return "users"
}
