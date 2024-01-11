package permission

import (
	"context"
	"github.com/pjimming/zustacm/core/internal/common"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRolePermissionTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRolePermissionTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolePermissionTreeLogic {
	return &GetRolePermissionTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRolePermissionTreeLogic) GetRolePermissionTree() (resp *types.GetResourceTreeResp, err error) {
	// todo: complete role

	resources := make([]*model.Resource, 0)
	if err = l.svcCtx.DB.Where("type = ?", "MENU").Find(&resources).Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resp = &types.GetResourceTreeResp{Resource: common.GetResourceTree(resources)}

	return
}
