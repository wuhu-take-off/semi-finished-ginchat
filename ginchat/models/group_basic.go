package models

import (
	"ginchat/utils"
	"gorm.io/gorm"
)

// 群信息
type GroupBasic struct {
	gorm.Model
	Name    string
	OwnerId uint
	Icon    string
	Type    int
	Desc    string
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
func InitGroupBasic() {
	err := utils.GetDB().AutoMigrate(GroupBasic{})
	if err != nil {
		panic(err)
	}
}
