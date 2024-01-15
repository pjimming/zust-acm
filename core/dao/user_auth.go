package dao

import (
	"github.com/pjimming/zustacm/core/model"
	"gorm.io/gorm"
)

var UserAuth = &userAuth{}

type userAuth struct {
}

func (*userAuth) FindOne(db *gorm.DB, id int64) (r *model.UserAuth, err error) {
	r = &model.UserAuth{}
	if err = db.Model(&model.UserAuth{}).
		Where("id = ?", id).
		First(r).
		Error; err != nil {
		return nil, err
	}
	return
}
