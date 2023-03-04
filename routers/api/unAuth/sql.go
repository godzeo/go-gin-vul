package unAuth

import (
	"github.com/gin-gonic/gin"
	"github.com/godzeo/go-gin-vul/service/safe_service"
	"github.com/godzeo/go-gin-vul/service/vul_service"
)

func Sqlli(c *gin.Context) {

	User := c.PostForm("username")
	Password := c.PostForm("password")
	//User := "test"

	println("Password=" + Password)

	// 自动创建一个表结构
	//var db *gorm.DB
	//db, _ = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	//	setting.DatabaseSetting.User,
	//	setting.DatabaseSetting.Password,
	//	setting.DatabaseSetting.Host,
	//	setting.DatabaseSetting.Name))
	//db.AutoMigrate(&Login{})

	loginService := vul_service.LogData{Username: User, Password: Password}
	isExist, err := loginService.LoginCheck()
	if err != nil {
		c.JSON(500, gin.H{
			"err": err,
		})
		return
	}

	if !isExist {
		c.JSON(403, gin.H{
			"success": "login fail",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": "login succeed " + User,
	})
}

func SqlliSafe(c *gin.Context) {

	User := c.PostForm("user")
	Password := c.PostForm("password")

	loginService := safe_service.LogData{Username: User, Password: Password}
	isExist, err := loginService.LoginCheck()
	if err != nil {
		c.JSON(500, gin.H{
			"err": err,
		})
		return
	}

	if !isExist {
		c.JSON(403, gin.H{
			"success": "login fail",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": "login succeed " + User,
	})
}
