package vul_service

import "github.com/godzeo/go-gin-vul/models"

type LogData struct {
	Username string
	Password string
}

func (a *LogData) LoginCheck() (bool, error) {
	return models.Slqlimode(a.Username, a.Password)
}
