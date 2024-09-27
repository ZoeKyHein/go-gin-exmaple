package v1

import (
	"github.com/ZoeKyHein/go-gin-example/models"
	"github.com/ZoeKyHein/go-gin-example/pkg/e"
	"github.com/ZoeKyHein/go-gin-example/pkg/setting"
	"github.com/ZoeKyHein/go-gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// AddTag 添加标签
func AddTag(c *gin.Context) {
	// todo 待实现
	name := c.Query("name") // 获取参数
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	// 校验参数
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistTagByName(name) {
			code = e.ERROR_EXIST_TAG
		} else {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// EditTag 编辑标签
func EditTag(c *gin.Context) {
	// todo 待实现
}

// DeleteTag 删除标签
func DeleteTag(c *gin.Context) {
	// todo 待实现
}

// GetTags 获取多个标签
func GetTags(c *gin.Context) {
	// todo 待实现
	name := c.Query("name") // 获取参数

	filterMaps := make(map[string]interface{}) // 存储过滤条件
	dataMaps := make(map[string]interface{})   // 存储数据

	// 筛选名称
	if name != "" {
		filterMaps["name"] = name
	}

	// 筛选状态
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		filterMaps["state"] = state
	}

	code := e.SUCCESS
	dataMaps["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, filterMaps)
	dataMaps["total"] = models.GetTagTotal(filterMaps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": dataMaps,
	})
}
