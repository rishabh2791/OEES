package persistance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"oees/domain/entity"
	"oees/domain/repository"
	"oees/domain/value_objects"
	"oees/infrastructure/utilities"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/crypto/bcrypt"
)

type authRepo struct {
	RedisClient *redis.Client
	Logger      hclog.Logger
	Conf        *utilities.ServerConfig
}

var _ repository.AuthRepository = &authRepo{}

func newAuthRepo(logger hclog.Logger, redisClient *redis.Client, conf *utilities.ServerConfig) *authRepo {
	authRepo := authRepo{}
	authRepo.RedisClient = redisClient
	authRepo.Logger = logger
	authRepo.Conf = conf
	return &authRepo
}

func (auth *authRepo) Authenticate(reqUser *entity.User, user *entity.User) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password))
}

func (auth *authRepo) GenerateTokens(user *entity.User) (*value_objects.Token, error) {
	token := value_objects.Token{}
	tokenData := auth.Conf.GetTokenConfig()
	keyData := auth.Conf.GetKeyConfig()
	token.ATExpires = time.Now().Add(time.Minute * time.Duration(tokenData.JWTExpiration)).Unix()
	token.AccessUUID = uuid.New().String()
	token.ATDuration = tokenData.JWTExpiration * 60
	token.RTExpires = time.Now().Add(time.Hour * 24 * time.Duration(tokenData.RefreshExpiration)).Unix()
	token.RefreshUUID = uuid.New().String()
	token.RTDuration = tokenData.RefreshExpiration * 24 * 60 * 60

	accessClaims := jwt.MapClaims{}
	accessClaims["username"] = user.Username
	accessClaims["token_type"] = "access"
	accessClaims["authorized"] = true
	accessClaims["duration"] = tokenData.JWTExpiration
	accessClaims["access_uuid"] = token.AccessUUID
	accessClaims["expiry"] = token.ATExpires

	atPrivateKey, err := ioutil.ReadFile(keyData.AccessTokenPrivateKeyPath)
	if err != nil {
		auth.Logger.Error("Unable to get Private Key String for Access Token.")
		return nil, err
	}

	atSigningKey, err := jwt.ParseRSAPrivateKeyFromPEM(atPrivateKey)
	if err != nil {
		auth.Logger.Error("Unable to get Signing Key String from Private Key for Access Token.")
		return nil, err
	}

	refreshClaims := jwt.MapClaims{}
	refreshClaims["username"] = user.Username
	refreshClaims["token_type"] = "refresh"
	refreshClaims["refresh_uuid"] = token.RefreshUUID
	refreshClaims["expiry"] = token.RTExpires
	refreshClaims["custom_key"] = auth.GenerateCustomKey(user.Username, user.SecretKey)

	rtPrivateKey, err := ioutil.ReadFile(keyData.RefreshTokenPrivateKeyPath)
	if err != nil {
		auth.Logger.Error("Unable to get Private Key String for Refresh Token.")
		return nil, err
	}

	rtSigningKey, err := jwt.ParseRSAPrivateKeyFromPEM(rtPrivateKey)
	if err != nil {
		auth.Logger.Error("Unable to get Signing Key String from Private String for Refresh Token.")
		return nil, err
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshClaims)
	token.Username = user.Username

	token.AccessToken, err = accessToken.SignedString(atSigningKey)
	if err != nil {
		auth.Logger.Error("Unable to sign Access Token.")
		return nil, err
	}
	token.RefreshToken, err = refreshToken.SignedString(rtSigningKey)
	if err != nil {
		auth.Logger.Error("Unable to sign Refresh Token.")
		return nil, err
	}
	return &token, nil
}

func (auth *authRepo) GenerateAuth(token *value_objects.Token) error {
	at := time.Unix(token.ATExpires, 0)
	rt := time.Unix(token.RTExpires, 0)
	now := time.Now()
	accessErr := auth.RedisClient.Set(token.AccessUUID, token.Username, at.Sub(now)).Err()
	if accessErr != nil {
		return accessErr
	}
	refreshErr := auth.RedisClient.Set(token.RefreshUUID, token.Username, rt.Sub(now)).Err()
	if refreshErr != nil {
		return refreshErr
	}
	return nil
}

func (auth *authRepo) GenerateCustomKey(username string, secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(username))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func (auth *authRepo) ValidateAccessToken(tokenString string) (*value_objects.AccessDetail, error) {
	accessDetails := value_objects.AccessDetail{}
	keyData := auth.Conf.GetKeyConfig()
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Unexpected Signing Method.")
		}

		publicKey, err := ioutil.ReadFile(keyData.AccessTokenPublicKeyPath)
		if err != nil {
			return nil, err
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
		if err != nil {
			return nil, err
		}
		return verifyKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["username"].(string) == "" || claims["token_type"].(string) != "access" {
		return nil, errors.New("Invalid Token. Authentication Failed.")
	}

	accessDetails.AccessUUID = claims["access_uuid"].(string)
	accessDetails.Username = claims["username"].(string)
	accessDetails.Duration = int(claims["duration"].(float64))

	return &accessDetails, nil
}

func (auth *authRepo) FetchAuth(uuid string) (string, error) {
	username, err := auth.RedisClient.Get(uuid).Result()
	if err != nil {
		return "", err
	}
	return username, nil
}

func (auth *authRepo) DeleteAuth(uuid string) (int64, error) {
	deleted, err := auth.RedisClient.Del(uuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func (auth *authRepo) ValidateRefreshToken(tokenString string) (*value_objects.RefreshDetail, error) {
	refreshDetails := value_objects.RefreshDetail{}
	keyData := auth.Conf.GetKeyConfig()
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Unexpected Signing Method.")
		}

		publicKey, err := ioutil.ReadFile(keyData.RefreshTokenPublicKeyPath)
		if err != nil {
			return nil, err
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
		if err != nil {
			return nil, err
		}

		return verifyKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["username"].(string) == "" || claims["token_type"].(string) != "refresh" {
		return nil, errors.New("")
	}

	refreshDetails.RefreshUUID = claims["refresh_uuid"].(string)
	refreshDetails.Username = claims["username"].(string)
	refreshDetails.CustomKey = claims["custom_key"].(string)

	return &refreshDetails, nil
}
