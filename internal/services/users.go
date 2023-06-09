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
	Get(userID uint) *postgres.User
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

func (s UserService) Get(userID uint) *postgres.User {
	return s.UserDAO.GetByID(userID)
}

func (s UserService) Authorize(username, password string) (string, error) {
	user := s.UserDAO.GetByUsername(username)

	if user == nil {
		return "", errors.New("user not found")
	}

	if !hash.MatchPasswords(password, user.Password) {
		return "", errors.New("passwords mismatch")
	}

	token, err := jwt.Encode(user.ID, s.Config.SecretKey)
	return token, err
}
