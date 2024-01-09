package auth

import (
	"context"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/utils/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuthCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthCaptchaLogic {
	return &GetAuthCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuthCaptchaLogic) GetAuthCaptcha() (resp *types.GetAuthCaptchaResp, err error) {
	captchaId, b64s, _, err := l.svcCtx.AuthCaptcha.Generate()
	if err != nil {
		err = errorx.Error500("generate captcha fail, %v")
		return nil, err
	}

	// logx.Debug(captchaId, b64s, answer)

	resp = &types.GetAuthCaptchaResp{
		CaptchaId: captchaId,
		B64s:      b64s,
	}

	return
}
