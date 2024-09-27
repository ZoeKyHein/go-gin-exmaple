package models

type Article struct {
	Model

	TagID int `json:"tag_id"` // 标签ID
	Tag   Tag `json:"tag"`    // 标签

	Title      string `json:"title"`       // 标题
	Desc       string `json:"desc"`        // 描述
	Content    string `json:"content"`     // 内容
	CreatedBy  string `json:"created_by"`  // 创建人
	ModifiedBy string `json:"modified_by"` // 修改人
	State      int    `json:"state"`       // 状态
}

// GetArticles 获取多个文章
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

// GetArticle 获取单个文章
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

// UpdateArticle 更新文章
func UpdateArticle(id int, data interface{}) error {
	return db.Model(&Article{}).Where("id = ?", id).Updates(data).Error
}

// AddArticle 添加文章
func AddArticle(data map[string]interface{}) error {
	article := Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	}

	return db.Create(&article).Error
}

// DeleteArticle 删除文章
func DeleteArticle(id int) error {
	return db.Where("id = ?", id).Delete(&Article{}).Error
}
