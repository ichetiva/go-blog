package services

import (
	"errors"

	"github.com/ichetiva/go-blog/config"
	"github.com/ichetiva/go-blog/internal/dao"
	"github.com/ichetiva/go-blog/pkg/hash"
	"github.com/ichetiva/go-blog/pkg/jwt"
	"github.com/ichetiva/go-blog/pkg/postgres"
	"gorm.io/gorm"
)

type UserService struct {
	UserDAO dao.IUserDAO
	Config  *config.Config
}

type IUserService interface {
	Create(username, password string) *postgres.User
	Authorize(username, password string) (string, error)
}

func NewUserService(db *gorm.DB, config *config.Config) IUserService {
	return UserService{
		UserDAO: &dao.UserDAO{DB: db},
		Config:  config,
	}
}

func (s UserService) Create(username, password string) *postgres.User {
	password, _ = hash.HashPassword(password)
	return s.UserDAO.Create(username, password)
}

func (s UserService) Authorize(username, password string) (string, error) {
	user := s.UserDAO.Get(username)

	if user == nil {
		return "", errors.New("user not found")
	}

	password, _ = hash.HashPassword(password)
	if user.Password != password {
		return "", errors.New("passwords mismatch")
	}

	token, err := jwt.NewJWT(user.ID, s.Config.SecretKey)
	return token, err
}
