package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

// BeforeCreate 创建前models callback
func (article *Article) BeforeCreate(scope *gorm.Scope) (err error) {
	return scope.SetColumn("CreatedOn", time.Now().Unix())
}

// BeforeUpdate 更新前models callback
func (article *Article) BeforeUpdate(scope *gorm.Scope) (err error) {
	return scope.SetColumn("ModifiedOn", time.Now().Unix())
}

// ExistArticleByName 根据名称判断标签是否存在
//func ExistArticleByName(name string) (bool, error) {
//	var article Article
//	if err := db.Select("id").Where("name = ?", name).First(&article).Error; err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return false, nil
//		}
//		return false, err
//	}
//	return true, nil
//}

// ExistArticleByID 根据ID判断标签是否存在
func ExistArticleByID(id int) (bool, error) {
	var article Article
	if err := db.Select("id").Where("id = ?", id).First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
