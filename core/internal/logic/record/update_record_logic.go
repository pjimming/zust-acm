package record

import (
	"context"

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
	// todo: add your logic here and delete this line

	return nil
}
