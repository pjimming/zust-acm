package auth

import (
	"context"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/userauth"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogoutLogic {
	return &AuthLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLogoutLogic) AuthLogout() error {
	// rm redis token
	token, ok := userauth.GetTokenFromCtx(l.ctx)
	if !ok {
		return errorx.ErrorTokenInvalid()
	}

	if err := l.svcCtx.Redis.Del(l.ctx, token).Err(); err != nil {
		return errorx.ErrorRedis(err)
	}

	return nil
}
