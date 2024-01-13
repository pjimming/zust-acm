package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/helper"
	"github.com/pjimming/zustacm/core/utils/sqlbuilder"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPageLogic {
	return &GetUserPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPageLogic) GetUserPage(req *types.GetUserPageReq) (resp *types.GetUserPageResp, err error) {

	resp = &types.GetUserPageResp{
		Items: make([]*types.User, 0),
		Total: 0,
	}

	enrollmentYear := helper.ConvertStringToInt64Array(req.EnrollmentYear)

	sb := sqlbuilder.NewSQLBuilder(l.svcCtx.DB.Model(&model.UserInfo{})).
		AndStringLike("username", req.Username).
		AndStringLike("cf_id", req.CfId).
		AndStringLike("atc_id", req.AtcId).
		AndStringInLike(req.Name, "name", "cname").
		AndIntEQ("gender", req.Gender).
		AndIntIn("enrollment_year", enrollmentYear)

	if err = sb.ToSession().Count(&resp.Total).Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	users := make([]*model.UserInfo, 0)
	if err = sb.Offset((req.Page - 1) * req.Size).
		Limit(req.Size).
		ToSession().
		Find(&users).
		Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	for _, user := range users {
		tmp := &types.User{}
		_ = copier.Copy(tmp, user)
		tmp.CreatedAt = helper.GetLocalDateTime(user.CreatedAt)
		tmp.UpdatedAt = helper.GetLocalDateTime(user.UpdatedAt)
		tmp.Roles = make([]*types.UserRole, 0)

		userRoleRels := make([]*model.UserRoleRel, 0)
		if err = l.svcCtx.DB.Model(&model.UserRoleRel{}).
			Select("role_id").
			Where("user_id = ?", user.ID).
			Find(&userRoleRels).
			Error; err != nil {
			err = errorx.ErrorDB(err)
			return nil, err
		}

		roleIds := make([]int64, 0)
		for _, urr := range userRoleRels {
			roleIds = append(roleIds, urr.RoleID)
		}

		roles := make([]*model.Role, 0)
		if err = l.svcCtx.DB.Model(&model.Role{}).
			Where("id in (?)", roleIds).
			Find(&roles).
			Error; err != nil {
			err = errorx.ErrorDB(err)
			return nil, err
		}

		_ = copier.Copy(&tmp.Roles, roles)

		resp.Items = append(resp.Items, tmp)
	}

	return
}
