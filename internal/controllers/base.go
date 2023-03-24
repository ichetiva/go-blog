package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ichetiva/go-blog/config"
	"gorm.io/gorm"
)

type BaseController struct {
	Config *config.Config
	DB     *gorm.DB
}

type Controller interface {
	Register(router *gin.Engine)
}
