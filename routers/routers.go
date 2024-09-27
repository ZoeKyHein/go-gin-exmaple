package routers

import (
	v1 "github.com/ZoeKyHein/go-gin-example/routers/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	tag := r.Group("/api/v1/tag")
	{
		tag.POST("/add", v1.AddTag)             // AddTag 添加标签
		tag.PUT("/edit/:id", v1.EditTag)        // EditTag 编辑标签
		tag.DELETE("/delete/:id", v1.DeleteTag) // DeleteTag 删除标签
		tag.GET("/get", v1.GetTags)             // GetTags 获取多个标签
	}
	return r
}
