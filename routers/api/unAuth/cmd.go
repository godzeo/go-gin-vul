package unAuth

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func CMD1(c *gin.Context) {

	var ipaddr string
	// Check the request method
	if c.Request.Method == "GET" {
		ipaddr = c.Query("ip")
	} else if c.Request.Method == "POST" {
		ipaddr = c.PostForm("ip")
	}

	Command := fmt.Sprintf("ping -c 4 %s", ipaddr)
	output, err := exec.Command("/bin/sh", "-c", Command).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"success": string(output),
	})
}

type MyMsg struct {
	Domain   string `json:"domain"`
	Password string `json:"password"`
}

func CMD2(c *gin.Context) {
	// ---> 声明结构体变量
	var a MyMsg
	// ---> 绑定数据
	if err := c.ShouldBindJSON(&a); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	output, _ := exec.Command("/bin/bash", "-c", "dig "+a.Domain).CombinedOutput() // python -c is also vulnerable
	println(output)
	// ---> 返回JSON的output 不进行base64
	outputdeBase64, _ := base64.StdEncoding.DecodeString(string(output))

	c.JSON(200, gin.H{
		"success": outputdeBase64,
	})
}
