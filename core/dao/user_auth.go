package dao

import (
	"gorm.io/gorm"

	"github.com/pjimming/zustacm/core/model"
)

var UserAuth = &userAuth{}

type userAuth struct {
}

func (*userAuth) FindOne(db *gorm.DB, id int64) (r *model.UserAuth, err error) {
	r = &model.UserAuth{}
	if err = db.Model(&model.UserAuth{}).
		Where("id = ?", id).
		First(r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*userAuth) Insert(db *gorm.DB, data *model.UserAuth) error {
	return db.Model(&model.UserAuth{}).
		Create(data).
		Error
}

func (*userAuth) UpdateOne(db *gorm.DB, data *model.UserAuth) error {
	return db.Model(&model.UserAuth{}).
		Where("id = ?", data.ID).
		Updates(data).
		Error
}

func (*userAuth) DeleteOne(db *gorm.DB, id int64) error {
	return db.Model(&model.UserAuth{}).
		Delete(&model.UserAuth{ID: id}).
		Error
}

// GetPage 上层传入sql，进行分页查询
func (*userAuth) GetPage(db *gorm.DB, page, size int) (r []*model.UserAuth, count int64, err error) {
	if err = db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	r = make([]*model.UserAuth, 0)
	if err = db.Offset((page - 1) * size).
		Limit(size).
		Find(&r).
		Error; err != nil {
		return nil, 0, err
	}
	return
}
