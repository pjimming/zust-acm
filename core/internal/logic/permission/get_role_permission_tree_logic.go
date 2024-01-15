package permission

import (
	"context"
	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/internal/common"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/sqlbuilder"
	"github.com/pjimming/zustacm/core/utils/userauth"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRolePermissionTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRolePermissionTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolePermissionTreeLogic {
	return &GetRolePermissionTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRolePermissionTreeLogic) GetRolePermissionTree() (resp *types.GetResourceTreeResp, err error) {

	uc, ok := userauth.GetUserFromCtx(l.ctx)
	if !ok {
		err = errorx.ErrorTokenInvalid()
		return nil, err
	}

	userInfo, err := dao.UserInfo.FindByUsername(l.svcCtx.DB, uc.Username)
	if err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	rRRels, err := dao.RoleResourceRel.SearchByRoleId(l.svcCtx.DB, userInfo.RoleID)
	if err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resourceIds := make([]int64, 0)
	for _, rRRel := range rRRels {
		resourceIds = append(resourceIds, rRRel.ResourceID)
	}

	sb := sqlbuilder.NewSQLBuilder(l.svcCtx.DB.Model(&model.Resource{})).
		AndStringEQ("type", "MENU")

	if !uc.IsAdmin {
		sb.AndIntIn("id", resourceIds)
	}

	resources := make([]*model.Resource, 0)
	if err = sb.ToSession().
		Find(&resources).
		Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resp = &types.GetResourceTreeResp{Resource: common.GetResourceTree(resources)}

	return
}
