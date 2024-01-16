package dao

import (
	"gorm.io/gorm"

	"github.com/pjimming/zustacm/core/model"
)

var Role = &role{}

type role struct {
}

func (*role) FindOne(db *gorm.DB, id int64) (r *model.Role, err error) {
	r = &model.Role{}
	if err = db.Model(&model.Role{}).
		Where("id = ?", id).
		First(r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*role) Insert(db *gorm.DB, data *model.Role) error {
	return db.Model(&model.Role{}).
		Create(data).
		Error
}

func (*role) UpdateOne(db *gorm.DB, data *model.Role) error {
	return db.Model(&model.Role{}).
		Where("id = ?", data.ID).
		Updates(data).
		Error
}

func (*role) DeleteOne(db *gorm.DB, id int64) error {
	return db.Model(&model.Role{}).
		Delete(&model.Role{ID: id}).
		Error
}

// GetPage 上层传入sql，进行分页查询
func (*role) GetPage(db *gorm.DB, page, size int) (r []*model.Role, count int64, err error) {
	if err = db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	r = make([]*model.Role, 0)
	if err = db.Offset((page - 1) * size).
		Limit(size).
		Find(&r).
		Error; err != nil {
		return nil, 0, err
	}
	return
}
