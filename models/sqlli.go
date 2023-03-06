package models

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

func Slqlimode(username, password string) (bool, error) {

	// 增
	//var data = Login{User: username, Password: password}
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
	//err := db.Order(password).First(&data).Error

	// user=user&password=123456 AND EXTRACTVALUE(9509,CONCAT(0x5c,(SELECT user from blog.blog_login LIMIT 0,1)))

	// 建立数据库连接
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/blog")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// 构建SQL查询语句
	sql := fmt.Sprintf("SELECT COUNT(*) FROM blog_auth WHERE username='%s' AND password='%s'", username, password)

	// 执行SQL查询语句
	var count int
	err = db.QueryRow(sql).Scan(&count)

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	//if data.ID > 0 {
	//	return true, nil
	//}
	if count == 1 {
		fmt.Print("Login succeeded")
		return true, nil
	} else {
		fmt.Println("Login failed")
		return false, nil
	}

}

func Slqlisafemode(username, password string) (bool, error) {
	var data = Logindata{Username: username, Password: password}

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

func SlqliByIDmode(userID string) (Sqliuserdata, error) {

	println("userID:" + userID)

	//db.Where("user = ?", "some_user").Find(&Logindata{})

	var qLogindata Logindata
	rawSQL := fmt.Sprintf("SELECT * FROM blog_auth WHERE id = '%s'", userID)

	if err := db.Raw(rawSQL).Scan(&qLogindata).Error; err != nil {
		return Sqliuserdata{RawSQL: rawSQL}, err
	}

	sqliUserdata := Sqliuserdata{
		Logindata: qLogindata,
		RawSQL:    rawSQL,
	}

	if qLogindata.ID == 0 {
		return sqliUserdata, fmt.Errorf("user not found")
	}

	return sqliUserdata, nil

}
