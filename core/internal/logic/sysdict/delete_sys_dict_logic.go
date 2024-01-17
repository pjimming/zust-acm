package sysdict

import (
	"context"

	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysDictLogic {
	return &DeleteSysDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysDictLogic) DeleteSysDict(req *types.DeleteSysDictReq) error {

	err := dao.SysDict.DeleteOne(l.svcCtx.DB, req.ID)
	if err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	return nil
}
