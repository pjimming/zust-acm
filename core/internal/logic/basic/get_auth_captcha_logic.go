package basic

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
	captcha, err := l.svcCtx.DigitCaptcha.DrawCaptcha("1234")
	if err != nil {
		err = errorx.Error500("generate captcha fail, %v")
		return nil, err
	}

	resp = &types.GetAuthCaptchaResp{Captcha: captcha.EncodeB64string()}
	return
}
