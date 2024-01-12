package resource

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddResourceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddResourceLogic {
	return &AddResourceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddResourceLogic) AddResource(req *types.AddResourceReq) (resp *types.AddResourceResp, err error) {

	resource := &model.Resource{}
	_ = copier.Copy(resource, req)

	if err = l.svcCtx.DB.Save(resource).Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resp = &types.AddResourceResp{ID: resource.ID}

	return
}
