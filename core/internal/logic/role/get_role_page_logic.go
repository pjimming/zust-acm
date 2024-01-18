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

	resp = &types.GetRolePageResp{
		Items: make([]*types.Role, 0),
		Total: 0,
	}

	sb := sqlbuilder.NewSQLBuilder(l.svcCtx.DB.Model(&model.Role{})).
		OrStringLike("name", req.Name).
		OrStringLike("code", req.Code)

	if err = sb.ToSession().Count(&resp.Total).Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	roles := make([]*model.Role, 0)
	if err = sb.Offset((req.Page - 1) * req.Size).
		Limit(req.Size).
		ToSession().
		Find(&roles).
		Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	for _, role := range roles {
		tmp := model2Types(role)
		tmp.ResourceIds = make([]int64, 0)

		roleResourceRels := make([]*model.RoleResourceRel, 0)
		if err = l.svcCtx.DB.Model(&model.RoleResourceRel{}).
			Select("resource_id").
			Where("role_id = ?", role.ID).
			Find(&roleResourceRels).
			Error; err != nil {
			err = errorx.ErrorDB(err)
			return nil, err
		}

		for _, rrr := range roleResourceRels {
			tmp.ResourceIds = append(tmp.ResourceIds, rrr.ResourceID)
		}

		resp.Items = append(resp.Items, tmp)
	}

	return
}
