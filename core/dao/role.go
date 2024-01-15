package dao

import (
	"github.com/pjimming/zustacm/core/model"
	"gorm.io/gorm"
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
