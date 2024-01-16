package record

import (
	"context"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecordLogic {
	return &GetRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRecordLogic) GetRecord(req *types.GetRecordReq) (resp *types.GetRecordResp, err error) {
	// todo: add your logic here and delete this line

	return
}
