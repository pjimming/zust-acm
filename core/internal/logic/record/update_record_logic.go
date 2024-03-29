package record

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/utils/errorx"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRecordLogic {
	return &UpdateRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRecordLogic) UpdateRecord(req *types.UpdateRecordReq) error {

	record, err := dao.Record.FindOne(l.svcCtx.DB, req.ID)
	if err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	_ = copier.Copy(record, req)
	if err = dao.Record.UpdateOne(l.svcCtx.DB, record); err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	return nil
}
