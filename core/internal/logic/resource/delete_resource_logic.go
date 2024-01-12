package resource

import (
	"context"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteResourceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteResourceLogic {
	return &DeleteResourceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteResourceLogic) DeleteResource(req *types.DeleteResourceReq) (resp *types.DeleteResourceResp, err error) {
	affect, err := l.DeepDelete(req.ID)
	if err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}
	resp = &types.DeleteResourceResp{DelectCount: affect}

	return
}

func (l *DeleteResourceLogic) DeepDelete(id int64) (affect int, err error) {
	resources := make([]*model.Resource, 0)
	if err = l.svcCtx.DB.Where("parent_id = ?", id).Find(&resources).Error; err != nil {
		return
	}

	affect = 0
	for _, item := range resources {
		count, err := l.DeepDelete(item.ID)
		if err != nil {
			return 0, err
		}
		affect += count
	}

	if err = l.svcCtx.DB.Delete(&model.Resource{ID: id}).Error; err != nil {
		return 0, err
	}
	affect++

	return
}
