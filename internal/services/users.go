package services

import (
	"github.com/ichetiva/go-blog/internal/dao"
	"github.com/ichetiva/go-blog/pkg/hash"
	"github.com/ichetiva/go-blog/pkg/postgres"
	"gorm.io/gorm"
)

type UserService struct {
	UserDAO dao.IUserDAO
}

type IUserService interface {
	Create(username, password string) *postgres.User
}

func NewUserService(db *gorm.DB) IUserService {
	return UserService{
		UserDAO: &dao.UserDAO{DB: db},
	}
}

func (s UserService) Create(username, password string) *postgres.User {
	password, _ = hash.HashPassword(password)
	return s.UserDAO.Create(username, password)
}
