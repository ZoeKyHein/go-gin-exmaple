package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Tag 标签模型
type Tag struct {
	Model // 继承模型

	Name       string `json:"name"`        // 标签名称
	State      int    `json:"state"`       // 标签状态
	CreatedBy  string `json:"created_by"`  // 创建人
	ModifiedBy string `json:"modified_by"` // 更新人
}

// BeforeCreate 创建前models callback
func (tag *Tag) BeforeCreate(scope *gorm.Scope) (err error) {
	return scope.SetColumn("CreatedOn", time.Now().Unix())
}

// BeforeUpdate 更新前models callback
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) (err error) {
	return scope.SetColumn("ModifiedOn", time.Now().Unix())
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

// AddTag 添加标签
func AddTag(name string, state int, createdBy string) error {
	return db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}).Error
}

// EditTag 编辑标签
func EditTag(id int, data interface{}) error {
	return db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error
}

// DeleteTag 删除标签
func DeleteTag(id int) error {
	return db.Where("id = ?", id).Delete(&Tag{}).Error
}
