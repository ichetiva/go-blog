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

func (c UserController) Register(router *gin.RouterGroup) {
	router.POST("/users/sign-up", c.SignUpView)
	router.POST("/users/sign-in", c.SignInView)
	router.GET("/users/me", c.GetMeView)
}

// @description Sign up
// @tags users
// @accept json
// @produce json
// @param username body string true "Username"
// @param password body string true "Password"
// @param password1 body string true "Repeat password"
// @router /users/sign-up [post]
func (c UserController) SignUpView(ctx *gin.Context) {
	var data schemes.ReqSignUp
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Non-valid data was received",
		})
		return
	}

	if data.Password != data.Password1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Passwords mismatch",
		})
		return
	}

	user := c.UserService.Create(data.Username, data.Password)
	if user == nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "User already exist",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// @description Sign in
// @tags users
// @accept json
// @produce json
// @param username body string true "Username"
// @param password body string true "Password"
// @router /users/sign-in [post]
func (c UserController) SignInView(ctx *gin.Context) {
	var data schemes.ReqSignIn
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Non-valid data was received",
		})
		return
	}

	token, err := c.UserService.Authorize(data.Username, data.Password)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": schemes.ResSignIn{
			AccessToken: token,
		},
	})
}

// @description Get me
// @tags users
// @accept json
// @produce json
// @param Token header string true "Access token"
// @router /users/me [get]
func (c UserController) GetMeView(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Token")
	claims, err := jwt.Decode(token, c.Config.SecretKey)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Invalid token",
		})
		return
	}

	user := c.UserService.Get(claims.UserID)
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
