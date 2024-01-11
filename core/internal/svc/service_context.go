package svc

import (
	"context"
	"github.com/pjimming/zustacm/core/internal/config"
	"github.com/redis/go-redis/v9"

	"github.com/mojocn/base64Captcha"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config      config.Config
	AuthCaptcha *base64Captcha.Captcha
	DB          *gorm.DB
	Redis       *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// database
	db, err := gorm.Open(mysql.Open(c.Mysql.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// redis
	rdc := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password,
	})

	if err = rdc.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
		AuthCaptcha: base64Captcha.NewCaptcha(
			base64Captcha.NewDriverDigit(40, 80, 4, 0.4, 15),
			base64Captcha.DefaultMemStore,
		),
		Redis: rdc,
	}
}
