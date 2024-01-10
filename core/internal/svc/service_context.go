package svc

import (
	"github.com/mojocn/base64Captcha"
	"github.com/pjimming/zustacm/core/internal/config"
	"github.com/pjimming/zustacm/core/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	AuthCaptcha   *base64Captcha.Captcha
	UserAuthModel model.UserAuthModel
	UserInfoModel model.UserInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DSN)

	return &ServiceContext{
		Config: c,
		AuthCaptcha: base64Captcha.NewCaptcha(
			base64Captcha.NewDriverDigit(40, 80, 4, 0.4, 15),
			base64Captcha.DefaultMemStore,
		),
		UserAuthModel: model.NewUserAuthModel(sqlConn),
		UserInfoModel: model.NewUserInfoModel(sqlConn),
	}
}
