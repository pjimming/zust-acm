package dao

import (
	"gorm.io/gorm"

	"github.com/pjimming/zustacm/core/model"
)

var Team = &team{}

type team struct {
}

func (*team) FindOne(db *gorm.DB, id int64) (r *model.Team, err error) {
	r = &model.Team{}
	if err = db.Model(&model.Team{}).
		Where("id = ?", id).
		First(r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*team) Insert(db *gorm.DB, data *model.Team) error {
	return db.Model(&model.Team{}).
		Create(data).
		Error
}

func (*team) UpdateOne(db *gorm.DB, data *model.Team) error {
	return db.Model(&model.Team{}).
		Where("id = ?", data.ID).
		Updates(data).
		Error
}

func (*team) DeleteOne(db *gorm.DB, id int64) error {
	return db.Model(&model.Team{}).
		Delete(&model.Team{ID: id}).
		Error
}

// GetPage 上层传入sql，进行分页查询
func (*team) GetPage(db *gorm.DB, page, size int) (r []*model.Team, count int64, err error) {
	if err = db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	r = make([]*model.Team, 0)
	if err = db.Offset((page - 1) * size).
		Limit(size).
		Find(&r).
		Error; err != nil {
		return nil, 0, err
	}
	return
}

func (*team) FindByIds(db *gorm.DB, ids []int64) (r []*model.Team, err error) {
	r = make([]*model.Team, 0)
	if err = db.Model(&model.Team{}).
		Where("id in (?)", ids).
		Find(&r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*team) FindAll(db *gorm.DB) (r []*model.Team, err error) {
	r = make([]*model.Team, 0)
	if err = db.Model(&model.Team{}).
		Find(&r).
		Error; err != nil {
		return nil, err
	}
	return
}
