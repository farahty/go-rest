package database

import "github.com/nimerfarahty/go-rest/models"

func Migrate() error {
	return Conn.AutoMigrate(
		&models.User{},
		&models.Todo{},
		&models.Post{},
	)
}
