package postgres

import (
	"fmt"

	"github.com/ichetiva/go-blog/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

type Post struct {
	gorm.Model
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
	Title   string
	Content string
}

func NewDB(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s?sslmode=disable",
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresHost,
		config.PostgresDatabase,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
