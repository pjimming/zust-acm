package dao

import (
	"gorm.io/gorm"

	"github.com/pjimming/zustacm/core/model"
)

var {{firstUpper .Model}} = &{{.Model}}{}

type {{.Model}} struct {
}

func (*{{.Model}}) FindOne(db *gorm.DB, id int64) (r *model.{{firstUpper .Model}}, err error) {
	r = &model.{{firstUpper .Model}}{}
	if err = db.Model(&model.{{firstUpper .Model}}{}).
		Where("id = ?", id).
		First(r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*{{.Model}}) Insert(db *gorm.DB, data *model.{{firstUpper .Model}}) error {
	return db.Model(&model.{{firstUpper .Model}}{}).
		Create(data).
		Error
}

func (*{{.Model}}) UpdateOne(db *gorm.DB, data *model.{{firstUpper .Model}}) error {
	return db.Model(&model.{{firstUpper .Model}}{}).
		Where("id = ?", data.ID).
		Updates(data).
		Error
}

func (*{{.Model}}) DeleteOne(db *gorm.DB, id int64) error {
	return db.Model(&model.{{firstUpper .Model}}{}).
		Delete(&model.{{firstUpper .Model}}{ID: id}).
		Error
}

// GetPage 上层传入sql，进行分页查询
func (*{{.Model}}) GetPage(db *gorm.DB, page, size int) (r []*model.{{firstUpper .Model}}, count int64, err error) {
	if err = db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	r = make([]*model.{{firstUpper .Model}}, 0)
	if err = db.Offset((page - 1) * size).
		Limit(size).
		Find(&r).
		Error; err != nil {
		return nil, 0, err
	}
	return
}
