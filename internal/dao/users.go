package dao

import (
	"github.com/ichetiva/go-blog/pkg/postgres"
	"gorm.io/gorm"
)

type UserDAO struct {
	DB *gorm.DB
}

type IUserDAO interface {
	Create(username, password string) *postgres.User
}

func (dao *UserDAO) Create(username, password string) *postgres.User {
	user := postgres.User{
		Username: username,
		Password: password,
	}
	dao.DB.Create(&user)
	return &user
}
