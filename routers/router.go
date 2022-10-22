package routers

import (
	"github.com/godzeo/go-gin-vul/routers/api/unAuth"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/godzeo/go-gin-vul/middleware/jwt"
	"github.com/godzeo/go-gin-vul/pkg/export"
	"github.com/godzeo/go-gin-vul/pkg/qrcode"
	"github.com/godzeo/go-gin-vul/pkg/upload"
	"github.com/godzeo/go-gin-vul/routers/api"
	"github.com/godzeo/go-gin-vul/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	// 未授权的漏洞
	apivul := r.Group("/api/vul")
	{
		apivul.POST("/sql/login", unAuth.Sqlli)
		apivul.POST("cmd1", unAuth.CMD1)
		apivul.POST("cmd2", unAuth.CMD2)
	}

	// 安全修复后
	apisafe := r.Group("/api/safe")
	{
		apisafe.POST("/sql/login", unAuth.SqlliSafe)
	}

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//导出标签
		r.POST("/tags/export", v1.ExportTag)
		//导入标签
		r.POST("/tags/import", v1.ImportTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//生成文章海报
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
