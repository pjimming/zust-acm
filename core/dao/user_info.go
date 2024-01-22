package dao

import (
	"gorm.io/gorm"

	"github.com/pjimming/zustacm/core/model"
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

func (*userInfo) Insert(db *gorm.DB, data *model.UserInfo) error {
	return db.Model(&model.UserInfo{}).
		Create(data).
		Error
}

func (*userInfo) UpdateOne(db *gorm.DB, data *model.UserInfo) error {
	return db.Model(&model.UserInfo{}).
		Where("id = ?", data.ID).
		Updates(data).
		Error
}

func (*userInfo) DeleteOne(db *gorm.DB, id int64) error {
	return db.Model(&model.UserInfo{}).
		Delete(&model.UserInfo{ID: id}).
		Error
}

// GetPage 上层传入sql，进行分页查询
func (*userInfo) GetPage(db *gorm.DB, page, size int) (r []*model.UserInfo, count int64, err error) {
	if err = db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	r = make([]*model.UserInfo, 0)
	if err = db.Offset((page - 1) * size).
		Limit(size).
		Find(&r).
		Error; err != nil {
		return nil, 0, err
	}
	return
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

func (*userInfo) FindAll(db *gorm.DB) (r []*model.UserInfo, err error) {
	r = make([]*model.UserInfo, 0)
	if err = db.Model(&model.UserInfo{}).
		Find(&r).
		Error; err != nil {
		return nil, err
	}
	return
}
