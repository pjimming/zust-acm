package sysdict

import (
	"context"

	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/sqlbuilder"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysDictAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysDictAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysDictAllLogic {
	return &GetSysDictAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysDictAllLogic) GetSysDictAll(req *types.GetSysDictReq) (resp *types.GetSysDictResp, err error) {

	resp = &types.GetSysDictResp{
		Items: make([]*types.SysDict, 0),
		Total: 0,
	}

	sb := sqlbuilder.NewSQLBuilder(l.svcCtx.DB.Model(&model.SysDict{})).
		AndStringLike("type", req.Type).
		AndStringLike("label", req.Label).
		AndStringLike("value", req.Value).
		OrderDesc("updated_at").
		ToSession()

	sysDicts, err := dao.SysDict.FindAll(sb)
	if err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	for _, sysDict := range sysDicts {
		resp.Items = append(resp.Items, modelToTypes(sysDict))
	}

	return
}
