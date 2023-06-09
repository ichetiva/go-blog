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
	GetByUsername(username string) *postgres.User
	GetByID(userID uint) *postgres.User
}

func (dao *UserDAO) Create(username, password string) *postgres.User {
	user := postgres.User{
		Username: username,
		Password: password,
	}
	result := dao.DB.Create(&user)
	if result.Error != nil {
		return nil
	}
	return &user
}

func (dao *UserDAO) GetByUsername(username string) *postgres.User {
	var user postgres.User
	dao.DB.Where("username = ?", username).First(&user)
	return &user
}

func (dao *UserDAO) GetByID(userID uint) *postgres.User {
	var user postgres.User
	dao.DB.Where("id = ?", userID).First(&user)
	return &user
}
