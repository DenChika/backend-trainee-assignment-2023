package services

import (
	"backend-trainee-assignment-2023/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	saltForHash     = "fjhsdjkhfnekj3h4j43443vhjsdfj"
	tokenExpiration = 12 * time.Hour
	tokenSigningKey = "arciyuinty37864bckjrh3jk2g"
)

type authorizationService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

func newAuthorizationService(repo repository.Authorization) *authorizationService {
	return &authorizationService{
		repo: repo,
	}
}

func (service *authorizationService) SignUp(username, password string) error {
	hash := service.generatePasswordHash(password)
	return service.repo.CreateUser(username, hash)
}

func (service *authorizationService) SignIn(username, password string) (string, error) {
	hash := service.generatePasswordHash(password)
	userId, err := service.repo.GetUser(username, hash)
	if err != nil {
		return "", err
	}
	claims := tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpiration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tokenSigningKey))
}

func (service *authorizationService) ParseToken(accessToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(tokenSigningKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	return claims.UserId, nil
}

func (service *authorizationService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(saltForHash)))
}
