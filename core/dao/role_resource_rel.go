package dao

import (
	"gorm.io/gorm"

	"github.com/pjimming/zustacm/core/model"
)

var RoleResourceRel = &roleResourceRel{}

type roleResourceRel struct {
}

func (*roleResourceRel) FindOne(db *gorm.DB, id int64) (r *model.RoleResourceRel, err error) {
	r = &model.RoleResourceRel{}
	if err = db.Model(&model.RoleResourceRel{}).
		Where("id = ?", id).
		First(r).
		Error; err != nil {
		return nil, err
	}
	return
}

func (*roleResourceRel) Insert(db *gorm.DB, data *model.RoleResourceRel) error {
	return db.Model(&model.RoleResourceRel{}).
		Create(data).
		Error
}

func (*roleResourceRel) UpdateOne(db *gorm.DB, data *model.RoleResourceRel) error {
	return db.Model(&model.RoleResourceRel{}).
		Where("id = ?", data.ID).
		Updates(data).
		Error
}

func (*roleResourceRel) DeleteOne(db *gorm.DB, id int64) error {
	return db.Model(&model.RoleResourceRel{}).
		Delete(&model.RoleResourceRel{ID: id}).
		Error
}

// GetPage 上层传入sql，进行分页查询
func (*roleResourceRel) GetPage(db *gorm.DB, page, size int) (r []*model.RoleResourceRel, count int64, err error) {
	if err = db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	r = make([]*model.RoleResourceRel, 0)
	if err = db.Offset((page - 1) * size).
		Limit(size).
		Find(&r).
		Error; err != nil {
		return nil, 0, err
	}
	return
}

func (*roleResourceRel) SearchByRoleId(db *gorm.DB, roleId int64) (r []*model.RoleResourceRel, err error) {
	r = make([]*model.RoleResourceRel, 0)
	if err = db.Model(&model.RoleResourceRel{}).
		Where("role_id = ?", roleId).
		Find(&r).
		Error; err != nil {
		return nil, err
	}
	return
}
