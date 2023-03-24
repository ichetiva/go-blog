package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ichetiva/go-blog/config"
	"github.com/ichetiva/go-blog/internal/schemes"
	"github.com/ichetiva/go-blog/internal/services"
	"github.com/ichetiva/go-blog/pkg/jwt"
	"gorm.io/gorm"
)

type UserController struct {
	BaseController
	UserService services.IUserService
}

func NewUserController(config *config.Config, db *gorm.DB) Controller {
	return UserController{
		BaseController: BaseController{
			Config: config,
			DB:     db,
		},
		UserService: services.NewUserService(db, config),
	}
}

func (c UserController) Register(router *gin.Engine) {
	router.POST("/users/sign-up", c.SignUpView)
	router.POST("/users/sign-in", c.SignInView)
}

func (c UserController) SignUpView(ctx *gin.Context) {
	var data schemes.ReqSignUp
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Non-valid data was received",
		})
	}

	if data.Password != data.Password1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Passwords not match",
		})
	}

	user := c.UserService.Create(data.Username, data.Password)
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (c UserController) SignInView(ctx *gin.Context) {
	var data schemes.ReqSignIn
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Non-valid data was received",
		})
	}

	token, err := c.UserService.Authorize(data.Username, data.Password)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": schemes.ResSignIn{
			AccessToken: token,
		},
	})
}

func (c UserController) GetMeView(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Token")
	claims, err := jwt.Decode(token, c.Config.SecretKey)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": err.Error(),
		})
	}

	user := c.UserService.Get(claims.UserID)
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
