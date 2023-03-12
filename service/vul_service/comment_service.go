package vul_service

import (
	"github.com/godzeo/go-gin-vul/models"
	"html/template"
)

// AddComment adds a new comment
func AddComment(username, content string) error {
	return models.CreateComment(template.HTML(username), template.HTML(content))
}

// GetComments retrieves all comments
func GetComments() ([]models.Comment, error) {
	return models.GetComments()
}
