package v1

import (
	"github.com/ZoeKyHein/go-gin-example/models"
	"github.com/ZoeKyHein/go-gin-example/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

// GetArticles 获取文章列表
func GetArticles(c *gin.Context) {
	dataMaps := make(map[string]interface{})   // 存储数据
	filterMaps := make(map[string]interface{}) // 存储过滤条件
	valid := validation.Validation{}

	// 校验状态
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		filterMaps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	// 校验标签
	var tagID int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagID = com.StrTo(arg).MustInt()
		filterMaps["tag_id"] = tagID
		valid.Min(tagID, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		dataMaps["tag_id"] = tagID
		dataMaps["state"] = state
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": dataMaps,
	})
}

// GetArticle 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if exist, err := models.ExistArticleByID(id); err != nil {
			code = e.ERROR
		} else if !exist {
			code = e.ERROR_NOT_EXIST_ARTICLE
		} else {
			data = models.GetArticle(id)
			code = e.SUCCESS
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddArticle 新增文章
func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("描述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("描述最长为255字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		data := make(map[string]interface{})
		data["tag_id"] = tagId
		data["title"] = title
		data["desc"] = desc
		data["content"] = content
		data["created_by"] = createdBy
		data["state"] = state

		code = e.SUCCESS
		if err := models.AddArticle(data); err != nil {
			code = e.ERROR
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// EditArticle 修改文章
func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}
	var state int = -1
	if args := c.Query("state"); args != "" {
		state = com.StrTo(args).MustInt()
		valid.Range(com.StrTo(args).MustInt(), 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Max(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("描述最长为255字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if exist, err := models.ExistArticleByID(id); err != nil {
			code = e.ERROR
		} else if !exist {
			code = e.ERROR_NOT_EXIST_ARTICLE
		} else {
			data := make(map[string]interface{})
			if tagId > 0 {
				data["tag_id"] = tagId
			}
			if title != "" {
				data["title"] = title
			}
			if desc != "" {
				data["desc"] = desc
			}
			if content != "" {
				data["content"] = content
			}
			data["modified_by"] = modifiedBy
			if state != -1 {
				data["state"] = state
			}

			if err := models.UpdateArticle(id, data); err != nil {
				code = e.ERROR
			}
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if exist, err := models.ExistArticleByID(id); err != nil {
			code = e.ERROR
		} else if !exist {
			code = e.ERROR_NOT_EXIST_ARTICLE
		} else {
			if err := models.DeleteArticle(id); err != nil {
				code = e.ERROR
			}
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
