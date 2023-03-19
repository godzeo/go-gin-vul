package unAuth

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/godzeo/go-gin-vul/models"
	service "github.com/godzeo/go-gin-vul/service/vul_service"
	"html/template"
	"net/http"
	"time"
)

func Xss(c *gin.Context) {
	var name string
	// Check the request method
	if c.Request.Method == "GET" {
		name = c.Query("name")
	} else if c.Request.Method == "POST" {
		name = c.PostForm("name")
	}

	html := fmt.Sprintf(`
		<h1>Hello, %s!</h1>
		<p>You entered the name: <strong>%s</strong></p>
		<a href="%s">This is a link %s</a>
		<iframe src="%s">%s</iframe>
		<span style="%s">This is a span %s</span>
	`, name, name, name, name, name, name, name, name)

	// 解析模板文件
	tmpl, err := template.ParseFiles("templates/display.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing template file: %v", err)
		return
	}

	// 合并模板和数据
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, gin.H{
		"Title": "XSS",
		"Body":  template.HTML(html),
	})
	if err != nil {
		c.String(http.StatusInternalServerError, "Error executing template: %v", err)
		return
	}

	// 将结果返回给浏览器
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(buf.Bytes())
}

// AddComment adds a new comment
func SafeAddComment(c *gin.Context) {

	var username string
	var content string
	// Check the request method
	if c.Request.Method == "GET" {
		username = c.Query("username")
		content = c.PostForm("content")
	} else if c.Request.Method == "POST" {
		username = c.PostForm("username")
		content = c.PostForm("content")
	}

	err := service.AddComment(username, content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "Comment added successfully",
	//})
	c.Redirect(http.StatusSeeOther, "/api/vul/getcomments")
}

// AddComment adds a new comment
func AddComment(c *gin.Context) {
	username := c.PostForm("username")
	content := c.PostForm("content")
	err := service.AddComment(username, content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "Comment added successfully",
	//})
	c.Redirect(http.StatusSeeOther, "/api/vul/getcomments")
}

// GetComments gets all comments
func SafeGetComments(c *gin.Context) {
	comments, err := service.GetComments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 构造数据，用于渲染模板
	data := struct {
		Title    string
		Year     int
		Comments []models.Comment
	}{
		Title:    "Comments",
		Year:     time.Now().Year(),
		Comments: comments,
	}

	// 解析模板
	tmpl, err := template.New("comment.html").ParseFiles("templates/comment.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 渲染模板
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回渲染后的 HTML
	c.Data(http.StatusOK, "text/html; charset=utf-8", buf.Bytes())
}

// GetComments gets all comments
func GetComments(c *gin.Context) {
	comments, err := service.GetComments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 将评论的内容转换为 template.HTML 类型
	for i, comment := range comments {
		comments[i].Content = template.HTML(comment.Content)
	}

	// 构造数据，用于渲染模板
	data := struct {
		Title    string
		Year     int
		Comments []models.Comment
	}{
		Title:    "Comments",
		Year:     time.Now().Year(),
		Comments: comments,
	}

	// 解析模板
	tmpl, err := template.New("comment.html").ParseFiles("templates/comment.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 渲染模板
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回渲染后的 HTML
	c.Data(http.StatusOK, "text/html; charset=utf-8", buf.Bytes())
}
