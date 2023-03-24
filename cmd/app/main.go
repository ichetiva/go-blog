package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ichetiva/go-blog/config"
	"github.com/ichetiva/go-blog/internal/controllers"
	"github.com/ichetiva/go-blog/pkg/postgres"
)

func main() {
	config := config.NewConfig()

	db, err := postgres.NewDB()
	if err != nil {
		log.Fatal("Connection to datbase error")
	}

	db.AutoMigrate(&postgres.User{})

	router := gin.Default()

	userController := controllers.NewUserController(config, db)
	userController.Register(router)

	router.Run()
}
