package services

import (
	"github.com/ichetiva/go-blog/internal/dao"
	"github.com/ichetiva/go-blog/pkg/postgres"
	"gorm.io/gorm"
)

type PostService struct {
	PostDAO dao.IPostDAO
	UserDAO dao.IUserDAO
}

type IPostService interface {
	Get(postID int) *postgres.Post
	Create(userID uint, title string, content string) *postgres.Post
}

func NewPostService(db *gorm.DB) IPostService {
	return PostService{
		PostDAO: &dao.PostDAO{DB: db},
		UserDAO: &dao.UserDAO{DB: db},
	}
}

func (s PostService) Get(postID int) *postgres.Post {
	return s.PostDAO.Get(postID)
}

func (s PostService) Create(userID uint, title string, content string) *postgres.Post {
	user := s.UserDAO.GetByID(userID)
	return s.PostDAO.Create(user, title, content)
}
