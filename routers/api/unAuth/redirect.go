package unAuth

import "github.com/gin-gonic/gin"

func Redirect(c *gin.Context) {
	loc := c.Query("redirect")
	c.Redirect(302, loc)
}

func SafeRedirect(c *gin.Context) {
	baseUrl := "https://baidu.com/path?q="
	loc := c.Query("redirect")
	c.Redirect(302, baseUrl+loc)
}
