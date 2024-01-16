package dao

import (
	"gorm.io/gorm"

	"github.com/pjimming/zustacm/core/model"
)

var Resource = &resource{}

type resource struct {
}

func (*resource) FindOne(db *gorm.DB, id int64) (r *model.Resource, err error) {
	r = &model.Resource{}
	if err = db.Model(&model.Resource{}).
		Where("id = ?", id).
		First(r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*resource) Insert(db *gorm.DB, data *model.Resource) error {
	return db.Model(&model.Resource{}).
		Create(data).
		Error
}

func (*resource) UpdateOne(db *gorm.DB, data *model.Resource) error {
	return db.Model(&model.Resource{}).
		Where("id = ?", data.ID).
		Updates(data).
		Error
}

func (*resource) DeleteOne(db *gorm.DB, id int64) error {
	return db.Model(&model.Resource{}).
		Delete(&model.Resource{ID: id}).
		Error
}

// GetPage 上层传入sql，进行分页查询
func (*resource) GetPage(db *gorm.DB, page, size int) (r []*model.Resource, count int64, err error) {
	if err = db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	r = make([]*model.Resource, 0)
	if err = db.Offset((page - 1) * size).
		Limit(size).
		Find(&r).
		Error; err != nil {
		return nil, 0, err
	}
	return
}
