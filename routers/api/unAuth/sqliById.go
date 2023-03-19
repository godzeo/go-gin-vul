package unAuth

import (
	"github.com/gin-gonic/gin"
	"github.com/godzeo/go-gin-vul/service/vul_service"
)

func SqlliById(c *gin.Context) {

	var userID string
	// Check the request method
	if c.Request.Method == "GET" {
		userID = c.Query("userid")
	} else if c.Request.Method == "POST" {
		userID = c.PostForm("userid")
	}

	println("userid=" + userID)

	queryService := vul_service.UserID{UserID: userID}

	QueryData, err := queryService.QueryByID()

	if err != nil {
		c.JSON(500, gin.H{
			"err": err,
		})
		return
	}

	if QueryData.Logindata.ID == 0 {
		c.JSON(403, gin.H{
			"err": "user not found",
		})
		return
	}

	c.JSON(200, QueryData)
}
