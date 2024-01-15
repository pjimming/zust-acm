package dao

import (
	"github.com/pjimming/zustacm/core/model"

	"gorm.io/gorm"
)

var UserInfo = &userInfo{}

type userInfo struct {
}

func (*userInfo) FindOne(db *gorm.DB, id int64) (r *model.UserInfo, err error) {
	r = &model.UserInfo{}
	if err = db.Model(&model.UserInfo{}).
		Where("id = ?", id).
		First(r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*userInfo) UpdateOne(db *gorm.DB, data *model.UserInfo) error {
	return db.Model(&model.UserInfo{}).
		Where("id = ?", data.ID).
		Updates(data).
		Error
}

func (*userInfo) FindByUsername(db *gorm.DB, username string) (r *model.UserInfo, err error) {
	r = &model.UserInfo{}
	if err = db.Model(&model.UserInfo{}).
		Where("username = ?", username).
		First(r).
		Error; err != nil {
		return nil, err
	}
	return
}
