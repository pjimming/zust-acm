package permission

import (
	"context"
	"github.com/jinzhu/copier"
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
	if err = l.svcCtx.DB.Find(&resources).Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resourcesMap := make(map[int64][]*model.Resource)
	for _, item := range resources {
		resourcesMap[item.ParentID] = append(resourcesMap[item.ParentID], item)
	}

	dummy := &types.Resource{ID: 0, Children: make([]*types.Resource, 0)}

	q := make([]*types.Resource, 0)
	q = append(q, dummy)

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		for _, resource := range resourcesMap[node.ID] {
			child := &types.Resource{}
			_ = copier.Copy(child, resource)
			child.Children = make([]*types.Resource, 0)
			node.Children = append(node.Children, child)
			q = append(q, child)
		}
	}

	resp = &types.GetResourceTreeResp{Resource: dummy.Children}

	return
}
