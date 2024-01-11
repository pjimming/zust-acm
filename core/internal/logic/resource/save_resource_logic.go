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

type SaveResourceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveResourceLogic {
	return &SaveResourceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveResourceLogic) SaveResource(req *types.SaveResourceReq) (resp *types.SaveResourceResp, err error) {

	resource := &model.Resource{ID: req.ID}
	if err = l.svcCtx.DB.First(resource).Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	_ = copier.Copy(resource, req)

	if err = l.svcCtx.DB.Model(resource).Save(resource).Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resp = &types.SaveResourceResp{ID: resource.ID}

	return
}
