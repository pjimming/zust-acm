package dao

import (
	"gorm.io/gorm"

	"github.com/pjimming/zustacm/core/model"
)

var Record = &record{}

type record struct {
}

func (*record) FindOne(db *gorm.DB, id int64) (r *model.Record, err error) {
	r = &model.Record{}
	if err = db.Model(&model.Record{}).
		Where("id = ?", id).
		First(r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*record) Insert(db *gorm.DB, data *model.Record) error {
	return db.Model(&model.Record{}).
		Create(data).
		Error
}

func (*record) UpdateOne(db *gorm.DB, data *model.Record) error {
	return db.Model(&model.Record{}).
		Where("id = ?", data.ID).
		Updates(data).
		Error
}

func (*record) DeleteOne(db *gorm.DB, id int64) error {
	return db.Model(&model.Record{}).
		Delete(&model.Record{ID: id}).
		Error
}

// GetPage 上层传入sql，进行分页查询
func (*record) GetPage(db *gorm.DB, page, size int) (r []*model.Record, count int64, err error) {
	if err = db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	r = make([]*model.Record, 0)
	if err = db.Offset((page - 1) * size).
		Limit(size).
		Find(&r).
		Error; err != nil {
		return nil, 0, err
	}
	return
}
