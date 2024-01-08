package svc

import (
	"github.com/mojocn/base64Captcha"
	"github.com/pjimming/zustacm/core/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	DigitCaptcha *base64Captcha.DriverDigit
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		DigitCaptcha: base64Captcha.NewDriverDigit(40, 80, 4, 0.4, 15),
	}
}
