package codeforces

import (
	"context"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SyncCfAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncCfAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncCfAllLogic {
	return &SyncCfAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncCfAllLogic) SyncCfAll() error {
	// todo: add your logic here and delete this line

	return nil
}
