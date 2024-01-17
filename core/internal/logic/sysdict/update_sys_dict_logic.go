package sysdict

import (
	"context"

	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysDictLogic {
	return &UpdateSysDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysDictLogic) UpdateSysDict(req *types.UpdateSysDictReq) error {

	sysDict, err := dao.SysDict.FindOne(l.svcCtx.DB, req.ID)
	if err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	_ = copier.Copy(sysDict, req)
	// todo: custom trans

	if err = dao.SysDict.UpdateOne(l.svcCtx.DB, sysDict); err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	return nil
}
