package repository

import (
	"oees/domain/entity"
	"oees/domain/value_objects"
)

type AuthRepository interface {
	Authenticate(reqUser *entity.User, user *entity.User) error
	GenerateTokens(user *entity.User) (*value_objects.Token, error)
	GenerateAuth(token *value_objects.Token) error
	GenerateCustomKey(username string, secretKey string) string
	ValidateAccessToken(token string) (*value_objects.AccessDetail, error)
	FetchAuth(uuid string) (string, error)
	DeleteAuth(uuid string) (int64, error)
	ValidateRefreshToken(token string) (*value_objects.RefreshDetail, error)
}
