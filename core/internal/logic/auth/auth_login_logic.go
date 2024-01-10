package auth

import (
	"context"
	"errors"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/helper"

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

	userAuth, err := l.svcCtx.UserAuthModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			err = errorx.Error400f("没有该[%s]用户", req.Username)
		} else {
			err = errorx.ErrorDB(err)
		}
		return nil, err
	}

	if err = helper.CheckPwd(userAuth.Password, req.Password); err != nil {
		err = errorx.ErrorAuth()
		return nil, err
	}

	resp = &types.AuthLoginResp{Token: "123"}

	return
}
