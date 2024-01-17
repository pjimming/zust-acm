package sysdict

import (
	"context"

	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/sqlbuilder"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysDictLogic {
	return &GetSysDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysDictLogic) GetSysDict(req *types.GetSysDictReq) (resp *types.GetSysDictResp, err error) {

	resp = &types.GetSysDictResp{
		Items: make([]*types.SysDict, 0),
		Total: 0,
	}

	sb := sqlbuilder.NewSQLBuilder(l.svcCtx.DB.Model(&model.SysDict{})).
		// todo: custom query
		OrderDesc("updated_at").
		ToSession()

	sysDicts, count, err := dao.SysDict.GetPage(sb, req.Page, req.Size)
	if err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resp.Total = count

	for _, sysDict := range sysDicts {
		resp.Items = append(resp.Items, modelToTypes(sysDict))
	}

	return
}
