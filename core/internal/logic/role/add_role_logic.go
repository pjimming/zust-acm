package role

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/sqlbuilder"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRoleLogic) AddRole(req *types.AddRoleReq) (resp *types.AddRoleResp, err error) {

	sb := sqlbuilder.NewSQLBuilder(l.svcCtx.DB.Model(&model.Role{})).
		OrStringEQ("code", req.Code).
		OrStringEQ("name", req.Name).
		ToSession()

	existCnt := int64(0)
	if err = sb.Count(&existCnt).Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	if existCnt > 0 {
		err = errorx.Error400f("角色名[%s]或角色编码[%s]已存在", req.Name, req.Code)
		return nil, err
	}

	role := &model.Role{}
	_ = copier.Copy(role, req)
	if err = l.svcCtx.DB.Create(role).Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resp = &types.AddRoleResp{ID: role.ID}

	roleResourceRels := make([]*model.RoleResourceRel, 0)
	for _, id := range req.ResourceIds {
		roleResourceRels = append(roleResourceRels, &model.RoleResourceRel{
			RoleID:     role.ID,
			ResourceID: id,
		})
	}

	if err = l.svcCtx.DB.Create(roleResourceRels).Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	return
}
