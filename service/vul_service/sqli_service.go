package vul_service

import "github.com/godzeo/go-gin-vul/models"

type LogData struct {
	Username string
	Password string
}
type UserID struct {
	UserID string
}

func (a *LogData) LoginCheck() (bool, error) {
	return models.Slqlimode(a.Username, a.Password)
}

func (b *UserID) QueryByID() (models.Sqliuserdata, error) {
	return models.SlqliByIDmode(b.UserID)
}
