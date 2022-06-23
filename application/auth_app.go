package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
	"oees/domain/value_objects"
)

type authApp struct {
	authRepository repository.AuthRepository
}

var _ AuthAppInterface = &authApp{}

func newAuthApp(authRepository repository.AuthRepository) *authApp {
	return &authApp{authRepository: authRepository}
}

func (auth *authApp) Authenticate(reqUser *entity.User, user *entity.User) error {
	return auth.authRepository.Authenticate(reqUser, user)
}

func (auth *authApp) GenerateTokens(user *entity.User) (*value_objects.Token, error) {
	return auth.authRepository.GenerateTokens(user)
}

func (auth *authApp) GenerateAuth(token *value_objects.Token) error {
	return auth.authRepository.GenerateAuth(token)
}

func (auth *authApp) GenerateCustomKey(username string, secretKey string) string {
	return auth.authRepository.GenerateCustomKey(username, secretKey)
}

func (auth *authApp) ValidateAccessToken(token string) (*value_objects.AccessDetail, error) {
	return auth.authRepository.ValidateAccessToken(token)
}

func (auth *authApp) FetchAuth(uuid string) (string, error) {
	return auth.authRepository.FetchAuth(uuid)
}

func (auth *authApp) DeleteAuth(uuid string) (int64, error) {
	return auth.authRepository.DeleteAuth(uuid)
}

func (auth *authApp) ValidateRefreshToken(token string) (*value_objects.RefreshDetail, error) {
	return auth.authRepository.ValidateRefreshToken(token)
}

type AuthAppInterface interface {
	Authenticate(reqUser *entity.User, user *entity.User) error
	GenerateTokens(user *entity.User) (*value_objects.Token, error)
	GenerateAuth(token *value_objects.Token) error
	GenerateCustomKey(username string, secretKey string) string
	ValidateAccessToken(token string) (*value_objects.AccessDetail, error)
	FetchAuth(uuid string) (string, error)
	DeleteAuth(uuid string) (int64, error)
	ValidateRefreshToken(token string) (*value_objects.RefreshDetail, error)
}
