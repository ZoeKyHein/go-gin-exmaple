package routers

import (
	"github.com/ZoeKyHein/go-gin-example/pkg/setting"
	v1 "github.com/ZoeKyHein/go-gin-example/routers/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	tag := r.Group("/api/v1")
	{
		tag.POST("/tags", v1.AddTag)          // AddTag 添加标签
		tag.PUT("/tags/:id", v1.EditTag)      // EditTag 编辑标签
		tag.DELETE("/tags/:id", v1.DeleteTag) // DeleteTag 删除标签
		tag.GET("/tags", v1.GetTags)          // GetTags 获取多个标签
	}

	article := r.Group("/api/v1")
	{
		article.POST("/articles", v1.AddArticle)          // AddArticle 添加文章
		article.PUT("/articles/:id", v1.EditArticle)      // EditArticle 编辑文章
		article.DELETE("/articles/:id", v1.DeleteArticle) // DeleteArticle 删除文章
		article.GET("/articles", v1.GetArticles)          // GetArticles 获取多个文章
	}
	return r
}
