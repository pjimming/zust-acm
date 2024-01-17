package record

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRecordLogic {
	return &AddRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRecordLogic) AddRecord(req *types.AddRecordReq) error {

	record := &model.Record{}
	_ = copier.Copy(record, req)

	if err := dao.Record.Insert(l.svcCtx.DB, record); err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	return nil
}
