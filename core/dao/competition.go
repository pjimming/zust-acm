package dao

import (
	"github.com/pjimming/zustacm/core/model"

	"gorm.io/gorm"
)

var Competition = &competition{}

type competition struct {
}

func (*competition) FindOne(db *gorm.DB, id int64) (r *model.Competition, err error) {
	r = &model.Competition{}
	if err = db.Model(&model.Competition{}).
		Where("id = ?", id).
		First(r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*competition) Insert(db *gorm.DB, data *model.Competition) error {
	return db.Model(&model.Competition{}).
		Create(data).
		Error
}

func (*competition) UpdateOne(db *gorm.DB, data *model.Competition) error {
	return db.Model(&model.Competition{}).
		Where("id = ?", data.ID).
		Updates(data).
		Error
}

func (*competition) DeleteOne(db *gorm.DB, id int64) error {
	return db.Model(&model.Competition{}).
		Delete(&model.Competition{ID: id}).
		Error
}

// GetPage 上层传入sql，进行分页查询
func (*competition) GetPage(db *gorm.DB, page, size int) (r []*model.Competition, count int64, err error) {
	if err = db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	r = make([]*model.Competition, 0)
	if err = db.Offset((page - 1) * size).
		Limit(size).
		Find(&r).
		Error; err != nil {
		return nil, 0, err
	}
	return
}

func (*competition) FindByIds(db *gorm.DB, ids []int64) (r []*model.Competition, err error) {
	r = make([]*model.Competition, 0)
	if err = db.Model(&model.Competition{}).
		Where("id in (?)", ids).
		Find(&r).
		Error; err != nil {
		return nil, err
	}
	return
}
