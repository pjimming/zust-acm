package resource

import (
	"context"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetResourceTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetResourceTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetResourceTreeLogic {
	return &GetResourceTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetResourceTreeLogic) GetResourceTree() (resp *types.GetResourceTreeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
