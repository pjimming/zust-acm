package dao

import (
	"gorm.io/gorm"

	"github.com/pjimming/zustacm/core/model"
)

var SysDict = &sysDict{}

type sysDict struct {
}

func (*sysDict) FindOne(db *gorm.DB, id int64) (r *model.SysDict, err error) {
	r = &model.SysDict{}
	if err = db.Model(&model.SysDict{}).
		Where("id = ?", id).
		First(r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*sysDict) Insert(db *gorm.DB, data *model.SysDict) error {
	return db.Model(&model.SysDict{}).
		Create(data).
		Error
}

func (*sysDict) UpdateOne(db *gorm.DB, data *model.SysDict) error {
	return db.Model(&model.SysDict{}).
		Where("id = ?", data.ID).
		Updates(data).
		Error
}

func (*sysDict) DeleteOne(db *gorm.DB, id int64) error {
	return db.Model(&model.SysDict{}).
		Delete(&model.SysDict{ID: id}).
		Error
}

// GetPage 上层传入sql，进行分页查询
func (*sysDict) GetPage(db *gorm.DB, page, size int) (r []*model.SysDict, count int64, err error) {
	if err = db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	r = make([]*model.SysDict, 0)
	if err = db.Offset((page - 1) * size).
		Limit(size).
		Find(&r).
		Error; err != nil {
		return nil, 0, err
	}
	return
}

func (*sysDict) FindByIds(db *gorm.DB, ids []int64) (r []*model.SysDict, err error) {
	r = make([]*model.SysDict, 0)
	if err = db.Model(&model.SysDict{}).
		Where("id in (?)", ids).
		Find(&r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*sysDict) FindByType(db *gorm.DB, typ string) (r []*model.SysDict, err error) {
	r = make([]*model.SysDict, 0)
	if err = db.Model(&model.SysDict{}).
		Where("type = ?", typ).
		Find(&r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*sysDict) FindAll(db *gorm.DB) (r []*model.SysDict, err error) {
	r = make([]*model.SysDict, 0)
	if err = db.Model(&model.SysDict{}).
		Find(&r).
		Error; err != nil {
		return nil, err
	}
	return
}
