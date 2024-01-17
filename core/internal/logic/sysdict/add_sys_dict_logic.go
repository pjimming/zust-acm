package sysdict

import (
	"context"

	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysDictLogic {
	return &AddSysDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysDictLogic) AddSysDict(req *types.AddSysDictReq) error {

	sysDict := &model.SysDict{}
	_ = copier.Copy(sysDict, req)
	// todo: custom trans

	if err := dao.SysDict.Insert(l.svcCtx.DB, sysDict); err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	return nil
}
