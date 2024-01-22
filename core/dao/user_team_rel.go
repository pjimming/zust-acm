package dao

import (
	"gorm.io/gorm"

	"github.com/pjimming/zustacm/core/model"
)

var UserTeamRel = &userTeamRel{}

type userTeamRel struct {
}

func (*userTeamRel) FindOne(db *gorm.DB, id int64) (r *model.UserTeamRel, err error) {
	r = &model.UserTeamRel{}
	if err = db.Model(&model.UserTeamRel{}).
		Where("id = ?", id).
		First(r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*userTeamRel) Insert(db *gorm.DB, data *model.UserTeamRel) error {
	return db.Model(&model.UserTeamRel{}).
		Create(data).
		Error
}

func (*userTeamRel) BatchAdd(db *gorm.DB, data []*model.UserTeamRel) error {
	if len(data) <= 0 {
		return nil
	}
	return db.Model(&model.UserTeamRel{}).
		Create(&data).
		Error
}

func (*userTeamRel) UpdateOne(db *gorm.DB, data *model.UserTeamRel) error {
	return db.Model(&model.UserTeamRel{}).
		Where("id = ?", data.ID).
		Updates(data).
		Error
}

func (*userTeamRel) DeleteOne(db *gorm.DB, id int64) error {
	return db.Model(&model.UserTeamRel{}).
		Delete(&model.UserTeamRel{ID: id}).
		Error
}

// GetPage 上层传入sql，进行分页查询
func (*userTeamRel) GetPage(db *gorm.DB, page, size int) (r []*model.UserTeamRel, count int64, err error) {
	if err = db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	r = make([]*model.UserTeamRel, 0)
	if err = db.Offset((page - 1) * size).
		Limit(size).
		Find(&r).
		Error; err != nil {
		return nil, 0, err
	}
	return
}

func (*userTeamRel) FindByIds(db *gorm.DB, ids []int64) (r []*model.UserTeamRel, err error) {
	r = make([]*model.UserTeamRel, 0)
	if err = db.Model(&model.UserTeamRel{}).
		Where("id in (?)", ids).
		Find(&r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*userTeamRel) FindByType(db *gorm.DB, typ string) (r []*model.UserTeamRel, err error) {
	r = make([]*model.UserTeamRel, 0)
	if err = db.Model(&model.UserTeamRel{}).
		Where("type = ?", typ).
		Find(&r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*userTeamRel) FindAll(db *gorm.DB) (r []*model.UserTeamRel, err error) {
	r = make([]*model.UserTeamRel, 0)
	if err = db.Model(&model.UserTeamRel{}).
		Find(&r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*userTeamRel) FindByUserIds(db *gorm.DB, ids []int64) (r []*model.UserTeamRel, err error) {
	r = make([]*model.UserTeamRel, 0)
	if err = db.Model(&model.UserTeamRel{}).
		Where("user_id in (?)", ids).
		Find(&r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*userTeamRel) FindByTeamIds(db *gorm.DB, ids []int64) (r []*model.UserTeamRel, err error) {
	r = make([]*model.UserTeamRel, 0)
	if err = db.Model(&model.UserTeamRel{}).
		Where("team_id in (?)", ids).
		Find(&r).
		Error; err != nil {
		return nil, err
	}
	return
}
