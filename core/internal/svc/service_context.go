package svc

import (
	"github.com/pjimming/zustacm/core/internal/config"

	"github.com/mojocn/base64Captcha"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config      config.Config
	AuthCaptcha *base64Captcha.Captcha
	DB          *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		AuthCaptcha: base64Captcha.NewCaptcha(
			base64Captcha.NewDriverDigit(40, 80, 4, 0.4, 15),
			base64Captcha.DefaultMemStore,
		),
		DB: db,
	}
}
