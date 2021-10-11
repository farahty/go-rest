package postsService

import (
	"github.com/nimerfarahty/go-rest/database"
	"github.com/nimerfarahty/go-rest/models"
)

func CreatePost(input models.CreatePostInput) (*models.Post, error) {

	post := &models.Post{
		Title:    input.Title,
		Body:     input.Body,
		AuthorID: *input.AuthorID,
	}

	tx := database.Conn.Create(&post)

	if tx.Error != nil {
		return nil, tx.Error
	}

	tx = database.Conn.Preload("Author").Find(&post)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return post, nil

}
