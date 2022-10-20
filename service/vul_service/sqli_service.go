package vul_service

import "github.com/EDDYCJY/go-gin-example/models"

type LogData struct {
	Username string
	Password string
}

func (a *LogData) LoginCheck() (bool, error) {
	return models.Slqlimode(a.Username, a.Password)
}
