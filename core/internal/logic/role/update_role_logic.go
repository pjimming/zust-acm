package role

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRoleReq) error {

	role, err := dao.Role.FindOne(l.svcCtx.DB, req.ID)
	if err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	_ = copier.Copy(role, req)
	// todo: custom trans

	if err = dao.Role.UpdateOne(l.svcCtx.DB, role); err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	if err = dao.RoleResourceRel.RemoveByRoleIdForce(l.svcCtx.DB, role.ID); err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	rRRels := make([]*model.RoleResourceRel, 0)
	for _, resourceId := range req.ResourceIds {
		rRRels = append(rRRels, &model.RoleResourceRel{
			RoleID:     role.ID,
			ResourceID: resourceId,
		})
	}

	if err = dao.RoleResourceRel.BatchAdd(l.svcCtx.DB, rRRels); err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	return nil
}
