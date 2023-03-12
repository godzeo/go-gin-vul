package models

import (
	"html/template"
	"time"
)

type Comment struct {
	ID        uint          `gorm:"primary_key" json:"id"`
	Username  string        `json:"username"`
	Content   template.HTML `json:"content"`
	CreatedAt time.Time     `json:"created_at"`
}

// CreateComment creates a new comment and saves it to the database
func CreateComment(username, content template.HTML) error {
	comment := Comment{
		Username:  string(username),
		Content:   content,
		CreatedAt: time.Now(),
	}
	return db.Create(&comment).Error
}

// GetComments retrieves all comments from the database
func GetComments() ([]Comment, error) {
	var comments []Comment
	err := db.Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
