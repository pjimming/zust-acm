package svc

import (
	"github.com/mojocn/base64Captcha"
	"github.com/pjimming/zustacm/core/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	AuthCaptcha *base64Captcha.Captcha
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		AuthCaptcha: base64Captcha.NewCaptcha(
			base64Captcha.NewDriverDigit(40, 80, 4, 0.4, 15),
			base64Captcha.DefaultMemStore,
		),
	}
}
