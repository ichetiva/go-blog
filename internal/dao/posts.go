package dao

import (
	"github.com/ichetiva/go-blog/pkg/postgres"
	"gorm.io/gorm"
)

type PostDAO struct {
	DB *gorm.DB
}

type IPostDAO interface {
	Get(postID int) *postgres.Post
	Create(user *postgres.User, title string, content string) *postgres.Post
}

func (dao *PostDAO) Get(postID int) *postgres.Post {
	var post postgres.Post
	dao.DB.Where("id = ?", postID).First(&post)
	return &post
}

func (dao *PostDAO) Create(user *postgres.User, title string, content string) *postgres.Post {
	post := postgres.Post{
		User:    *user,
		Title:   title,
		Content: content,
	}
	dao.DB.Create(&post)
	return &post
}
