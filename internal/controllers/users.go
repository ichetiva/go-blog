package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ichetiva/go-blog/config"
	"github.com/ichetiva/go-blog/internal/schemes"
	"github.com/ichetiva/go-blog/internal/services"
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
		UserService: services.NewUserService(db),
	}
}

func (c UserController) Register(router *gin.Engine) {
	router.POST("/users/sign-up", c.SignUpView)
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
