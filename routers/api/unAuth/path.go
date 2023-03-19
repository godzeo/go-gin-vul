package unAuth

import (
	"archive/zip"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func FileRead(c *gin.Context) {

	var path string

	// Check the request method
	if c.Request.Method == "GET" {
		path = c.Query("filename")
	} else if c.Request.Method == "POST" {
		path = c.PostForm("filename")
	}
	// Unfiltered file paths
	data, _ := ioutil.ReadFile(path)

	c.JSON(200, gin.H{
		"success": "read: " + string(data),
	})

}

func DirFile(c *gin.Context) {

	var dir string

	// Check the request method
	if c.Request.Method == "GET" {
		dir = c.Query("filename")
	} else if c.Request.Method == "POST" {
		dir = c.PostForm("filename")
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid directory",
		})
		return
	}

	var fileNames []string
	for _, file := range files {
		if file.IsDir() {
			fileNames = append(fileNames, file.Name()+"/")
		} else {
			fileNames = append(fileNames, file.Name())
		}
	}

	c.JSON(200, gin.H{
		"files": fileNames,
	})
}

// arbitrary file remove
func Fileremove(c *gin.Context) {
	var path string

	// Check the request method
	if c.Request.Method == "GET" {
		path = c.Query("filename")
	} else if c.Request.Method == "POST" {
		path = c.PostForm("filename")
	}
	os.Remove(path)
}

var content = ""

// bad: arbitrary file write
func Unzip(c *gin.Context) {

	var path string
	var text string

	// Check the request method
	if c.Request.Method == "GET" {
		path = c.Query("filename")
		text = c.Query("text")
	} else if c.Request.Method == "POST" {
		path = c.PostForm("filename")
		text = c.PostForm("text")
	}

	file_path := filepath.Join("/bin/", path)
	r, _ := zip.OpenReader(file_path)

	var abspath string
	for _, f := range r.File {
		abspath, _ = filepath.Abs(f.Name)
		ioutil.WriteFile(abspath, []byte(text), 0640)
	}

	data, _ := ioutil.ReadFile(abspath)

	c.JSON(200, gin.H{
		"success": "read: " + string(data),
	})
}

// safe fix
func Unzipsafe(c *gin.Context) {
	var path string

	// Check the request method
	if c.Request.Method == "GET" {
		path = c.Query("filename")
	} else if c.Request.Method == "POST" {
		path = c.PostForm("filename")
	}

	file_path := filepath.Join("/bin/", path)
	r, err := zip.OpenReader(file_path)
	if err != nil {
		fmt.Println("read zip file fail")
		c.JSON(500, gin.H{
			"success": "err: " + err.Error(),
		})
	}
	for _, f := range r.File {
		if !strings.Contains(f.Name, "..") {
			p, _ := filepath.Abs(f.Name)
			ioutil.WriteFile(p, []byte("present"), 0640)
		} else {
			c.JSON(500, gin.H{
				"success": "err: " + err.Error(),
			})
		}
	}
	c.JSON(200, gin.H{
		"success": "1: ",
	})
}
