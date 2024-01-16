package record

import (
	"context"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRecordLogic {
	return &DeleteRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRecordLogic) DeleteRecord(req *types.DeleteRecordReq) error {
	// todo: add your logic here and delete this line

	return nil
}
