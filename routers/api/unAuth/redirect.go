package unAuth

import "github.com/gin-gonic/gin"

func Redirect(c *gin.Context) {
	var loc string
	// Check the request method
	if c.Request.Method == "GET" {
		loc = c.Query("redirect")
	} else if c.Request.Method == "POST" {
		loc = c.PostForm("redirect")
	}
	c.Redirect(302, loc)
}

func SafeRedirect(c *gin.Context) {
	baseUrl := "https://baidu.com/path?q="
	loc := c.Query("redirect")
	c.Redirect(302, baseUrl+loc)
}
