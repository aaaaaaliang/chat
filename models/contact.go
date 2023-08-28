package models

import (
	"chat/utils"
	"fmt"
	"gorm.io/gorm"
)

// 人员关系
type Contact struct {
	gorm.Model
	OwnerId  uint //谁的关系信息
	TargetID uint //对应的谁
	Type     int  //对应的类型1好友 2群主 3
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}

func SearchFriend(userID int) []UserBasic {
	contacts := make([]Contact, 0)
	//存储目标id的用户
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id = ? and type=1", userID).Find(&contacts)
	for _, v := range contacts {
		fmt.Println(">>>>>>>>>>>>>", v)
		// 将目标用户的 ID 添加到 objIds 切片中
		objIds = append(objIds, uint64(v.TargetID))
	}
	// 创建一个空的 UserBasic 切片来存储查询结果
	users := make([]UserBasic, 0)
	utils.DB.Where("id in ?", objIds).Find(&users)
	return users
}
