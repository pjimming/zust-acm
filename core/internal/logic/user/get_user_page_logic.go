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

	sb := sqlbuilder.NewSQLBuilder(l.svcCtx.DB.Model(&model.UserInfo{})).
		AndStringLike("username", req.Username).
		AndStringLike("cf_id", req.CfId).
		AndStringLike("atc_id", req.AtcId).
		AndStringInLike(req.Name, "name", "cname").
		AndIntEQ("gender", req.Gender).
		AndIntIn("enrollment_year", req.EnrollmentYear)

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
		tmp.Role = &types.Role{}

		role := &model.Role{}
		if err = l.svcCtx.DB.Model(&model.Role{}).
			Where("id = ?", user.RoleID).
			First(&role).
			Error; err != nil {
			err = errorx.ErrorDB(err)
			return nil, err
		}

		_ = copier.Copy(tmp.Role, role)
		tmp.Role.CreatedAt = helper.GetLocalDateTime(role.CreatedAt)
		tmp.Role.UpdatedAt = helper.GetLocalDateTime(role.UpdatedAt)

		resp.Items = append(resp.Items, tmp)
	}

	return
}
