package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

// ExistTagByName 根据名称判断标签是否存在
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	if err := db.Select("id").Where("name = ?", name).First(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// ExistTagByID 根据ID判断标签是否存在
func ExistTagByID(id int) (bool, error) {
	var tag Tag
	if err := db.Select("id").Where("id = ?", id).First(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
