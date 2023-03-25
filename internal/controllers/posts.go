package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ichetiva/go-blog/config"
	"github.com/ichetiva/go-blog/internal/schemes"
	"github.com/ichetiva/go-blog/internal/services"
	"github.com/ichetiva/go-blog/pkg/jwt"
	"gorm.io/gorm"
)

type PostController struct {
	BaseController
	PostService services.IPostService
}

func NewPostController(config *config.Config, db *gorm.DB) Controller {
	return PostController{
		BaseController: BaseController{
			Config: config,
			DB:     db,
		},
		PostService: services.NewPostService(db),
	}
}

func (c PostController) Register(router *gin.Engine) {
	router.GET("/posts/get", c.GetPostView)
	router.POST("/posts/create", c.CreatePostView)
}

func (c PostController) GetPostView(ctx *gin.Context) {
	postIDString, ok := ctx.GetQuery("postId")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Require query param 'postId'",
		})
		return
	}

	postID, err := strconv.Atoi(postIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Non-valid post ID",
		})
		return
	}

	post := c.PostService.Get(postID)
	if post == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func (c PostController) CreatePostView(ctx *gin.Context) {
	var data schemes.ReqCreatePost
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Non-valid data was received",
		})
	}

	token := ctx.Request.Header.Get("Token")
	claims, err := jwt.Decode(token, c.Config.SecretKey)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Invalid token",
		})
		return
	}

	post := c.PostService.Create(claims.UserID, data.Title, data.Content)
	ctx.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}
