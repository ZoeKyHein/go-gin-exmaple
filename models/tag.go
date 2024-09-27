package models

// Tag 标签模型
type Tag struct {
	Model // 继承模型

	Name       string `json:"name"`        // 标签名称
	State      int    `json:"state"`       // 标签状态
	CreatedBy  string `json:"created_by"`  // 创建人
	ModifiedBy string `json:"modified_by"` // 更新人
}

// GetTags 获取标签列表
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

// GetTagTotal 获取标签总数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

// ExistTagByName 根据名称判断标签是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// AddTag 添加标签
func AddTag(name string, state int, createdBy string) error {
	return db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}).Error
}
