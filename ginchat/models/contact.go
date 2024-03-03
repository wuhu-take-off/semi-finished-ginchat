package models

import (
	"ginchat/utils"
	"gorm.io/gorm"
)

// 人员关系
type Contact struct {
	gorm.Model
	OwnerId  uint //谁的关系信息
	TargetId uint //对应的谁
	Type     int  //关系类型
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}

func InitContact() {
	err := utils.GetDB().AutoMigrate(Contact{})
	if err != nil {
		panic(err)
	}
}
