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
	path := c.Query("filename")

	// Unfiltered file paths
	data, _ := ioutil.ReadFile(path)

	c.JSON(200, gin.H{
		"success": "read: " + string(data),
	})

}

func Dirfile(c *gin.Context) {
	path := c.Query("filename")
	data, _ := ioutil.ReadFile(filepath.Join("/usr", path))

	c.JSON(200, gin.H{
		"success": "read: " + string(data),
	})

}

// arbitrary file remove
func Fileremove(c *gin.Context) {
	path := c.Query("path")
	os.Remove(path)
}

var content = ""

// bad: arbitrary file write
func Unzip(c *gin.Context) {
	path := c.Query("filename")
	text := c.Query("text")
	file_path := filepath.Join("/Users/zy/", path)
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
	path := c.Query("filename")
	file_path := filepath.Join("/Users/zy/", path)
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
