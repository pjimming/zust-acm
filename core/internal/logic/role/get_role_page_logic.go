package role

import (
	"context"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRolePageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRolePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolePageLogic {
	return &GetRolePageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRolePageLogic) GetRolePage(req *types.GetRolePageReq) (resp *types.GetRolePageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
