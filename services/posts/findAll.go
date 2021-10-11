package postsService

import (
	"github.com/nimerfarahty/go-rest/database"
	"github.com/nimerfarahty/go-rest/models"
)

func FindAll() ([]*models.Post, error) {

	posts := []*models.Post{}

	tx := database.Conn.Preload("Author").Find(&posts)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return posts, nil

}
