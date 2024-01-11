package resource

import (
	"context"

	"github.com/pjimming/zustacm/core/internal/common"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"

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

	resources := make([]*model.Resource, 0)
	if err = l.svcCtx.DB.Where("type = ?", "MENU").Find(&resources).Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resp = &types.GetResourceTreeResp{Resource: common.GetResourceTree(resources)}

	return
}
