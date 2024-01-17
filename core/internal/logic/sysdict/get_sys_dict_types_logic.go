package sysdict

import (
	"context"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysDictTypesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysDictTypesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysDictTypesLogic {
	return &GetSysDictTypesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysDictTypesLogic) GetSysDictTypes() (resp *types.GetSysDictTypesResp, err error) {

	resp = &types.GetSysDictTypesResp{Types: make([]string, 0)}

	sysDicts := make([]*model.SysDict, 0)
	if err = l.svcCtx.DB.Model(&model.SysDict{}).
		Distinct("type").
		Find(&sysDicts).
		Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	for _, sysDict := range sysDicts {
		resp.Types = append(resp.Types, sysDict.Type)
	}

	return
}
