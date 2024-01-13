package user

import (
	"context"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignRolesLogic {
	return &AssignRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignRolesLogic) AssignRoles(req *types.AssignRolesReq) (err error) {
	if err = l.svcCtx.DB.Model(&model.UserRoleRel{}).
		Where("user_id = ?", req.ID).
		Unscoped().
		Delete(&model.UserRoleRel{}).
		Error; err != nil {
		err = errorx.ErrorDB(err)
		return
	}

	urr := make([]*model.UserRoleRel, 0)
	for _, roleId := range req.RoleIds {
		urr = append(urr, &model.UserRoleRel{
			UserID: req.ID,
			RoleID: roleId,
		})
	}

	if err = l.svcCtx.DB.Create(urr).Error; err != nil {
		err = errorx.ErrorDB(err)
		return
	}

	return nil
}
