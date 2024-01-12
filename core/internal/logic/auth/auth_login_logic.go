package auth

import (
	"context"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/constant"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/userauth"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLoginLogic {
	return &AuthLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLoginLogic) AuthLogin(req *types.AuthLoginReq) (resp *types.AuthLoginResp, err error) {

	if !l.svcCtx.AuthCaptcha.Verify(req.CaptchaId, req.Captcha, true) {
		err = errorx.ErrorAuthCaptcha()
		return nil, err
	}

	userAuth := &model.UserAuth{}
	err = l.svcCtx.DB.Model(&model.UserAuth{}).Where("username = ?", req.Username).Limit(1).Find(userAuth).Error
	if err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	if err = userauth.CheckPwd(userAuth.Password, req.Password); err != nil {
		err = errorx.ErrorAuth()
		return nil, err
	}

	token, err := userauth.GenToken(req.Username)
	if err != nil {
		err = errorx.Error500f("gen token fail, %v", err)
		return nil, err
	}
	resp = &types.AuthLoginResp{Token: token}

	// set token
	if err = l.svcCtx.Redis.Set(l.ctx, userauth.GenTokenKey(token), 1, constant.TokenExpire).
		Err(); err != nil {
		err = errorx.ErrorRedis(err)
		return nil, err
	}

	return
}
