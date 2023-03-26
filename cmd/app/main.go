package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ichetiva/go-blog/config"
	"github.com/ichetiva/go-blog/docs"
	"github.com/ichetiva/go-blog/internal/controllers"
	"github.com/ichetiva/go-blog/pkg/postgres"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Blog API
// @version 1.0
func main() {
	config := config.NewConfig()

	db, err := postgres.NewDB(config)
	if err != nil {
		log.Fatal("Connection to datbase error")
	}

	db.AutoMigrate(&postgres.User{})
	db.AutoMigrate(&postgres.Post{})

	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := router.Group("/api")
	{
		userController := controllers.NewUserController(config, db)
		userController.Register(api)

		postController := controllers.NewPostController(config, db)
		postController.Register(api)
	}

	router.Run()
}
