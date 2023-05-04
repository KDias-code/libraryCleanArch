package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	// "hash"

	todo "github.com/KDias-code/todoapp"
	"github.com/KDias-code/todoapp/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt     = "asdardirqweo231koads"
	signKey  = "fjdisi321ewqieqw"
	tokenTTl = 12 * time.Hour
)

type AuthService struct {
	repo repository.Autharization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json: "user_id"`
}

func NewAuthService(repo repository.Autharization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid singing method")
		}

		return []byte(signKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GetAll(userId int) ([]todo.User, error) {
	return s.repo.GetAll(userId)
}

func (s *AuthService) Delete(userId int) error {
	return s.repo.Delete(userId)
}

func (s *AuthService) Update(userId, authorId int, input todo.UpdateUserInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, authorId, input)
}
