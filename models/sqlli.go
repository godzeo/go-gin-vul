package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Login struct {
	ID       uint   `gorm:"primary_key"`
	User     string `form:"user" json:"user" xml:"user"  binding:"required,userValidation"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func Slqlimode(username, password string) (bool, error) {

	// 增
	var data = Login{User: username, Password: password}
	//err2 := db.Create(&data)

	//正常情况
	//err := db.Select("id").Where(&data).First(&logindata).Error

	// 常规问题 字符串拼接
	// 约定：username为用户可控的数据，比如可以是req.user的值
	//err := db.Select(username).First(&data).Error    // user=(updatexml(1,concat(0x7e,(select user()),0x7e),1))&password=123456
	//err := db.Where(fmt.Sprintf("name = '%s'", username)).Find(&data).Error
	//err := db.Model(&data).Pluck(username, &username).Error
	//err := db.Group(User).First(&data).Error
	// err := db.Group("User").Having(username).First(&data).Error
	//err := db.Exec("select User from blog.blog_login where User = " + username).First(&data).Error //user=(updatexml(1,concat(0x7e,(select user()),0x7e),1))&password=123456
	//err := db.Raw("select User from blog.blog_login where User ="+ username).First(&data).Error

	// SELECT * FROM `blog`.`blog_login` ORDER BY `id` LIMIT 0,1000
	err := db.Order(password).First(&data).Error

	// user=user&password=123456 AND EXTRACTVALUE(9509,CONCAT(0x5c,(SELECT user from blog.blog_login LIMIT 0,1)))

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if data.ID > 0 {
		return true, nil
	}

	return false, nil
}

func Slqlisafemode(username, password string) (bool, error) {
	var data = Login{User: username, Password: password}

	// 对于表名
	validCols := map[string]bool{"user": true, "password": true}

	if _, ok := validCols[password]; !ok {
		fmt.Println("illegal column")
		return false, nil
	}
	err := db.Order(password).First(&data).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if data.ID > 0 {
		return true, nil
	}

	return false, nil
}
