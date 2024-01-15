package dao

import (
	"github.com/pjimming/zustacm/core/model"
	"gorm.io/gorm"
)

var RoleResourceRel = &roleResourceRel{}

type roleResourceRel struct {
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
