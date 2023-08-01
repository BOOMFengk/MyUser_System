package dao

import (
	"MyUser_System/internal/model"
	"MyUser_System/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// GetUserByName
//
//	@Description: 根据姓名获取用户
//	@param name
//	@return *model.User
//	@return error
func GetUserByName(name string) (*model.User, error) {
	user := &model.User{}
	if err := utils.GetDB().Model(model.User{}).Where("name=?", name).First(user).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		log.Errorf("GetUserByName fail:%v", err)
		return nil, fmt.Errorf("GetUsetByName fail:%v", err)
	}
	return user, nil

}

// CreateUser
//
//	@Description: 创建一个用户
//	@param user
//	@return error
func CreateUser(user *model.User) error {
	if err := utils.GetDB().Model(&model.User{}).Create(user).Error; err != nil {
		log.Errorf("CreateUser fail: %v", err)
		return fmt.Errorf("CreateUser fail:%v", err)

	}
	log.Infof("insert success")
	return nil

}

// UpdateUserInfo
//
//	@Description:更新昵称
func UpdateUserInfo(userName string, user *model.User) int64 {
	return utils.GetDB().Model(&model.User{}).Where("`name`=?", userName).Updates(user).RowsAffected
}
