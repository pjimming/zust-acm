package role

import (
	"context"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/sqlbuilder"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleAllLogic {
	return &GetRoleAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleAllLogic) GetRoleAll(req *types.GetRoleAllReq) (resp *types.GetRoleAllResp, err error) {

	resp = &types.GetRoleAllResp{Items: make([]*types.Role, 0)}

	sb := sqlbuilder.NewSQLBuilder(l.svcCtx.DB.Model(&model.Role{})).
		OrStringLike("name", req.Name).
		OrStringLike("code", req.Code)

	if req.IsEnable {
		sb.AndIntEQ("is_enable", 1)
	}

	roles := make([]*model.Role, 0)
	if err = sb.ToSession().Find(&roles).Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	for _, role := range roles {
		tmp := model2Types(role)
		resp.Items = append(resp.Items, tmp)
	}

	return
}
