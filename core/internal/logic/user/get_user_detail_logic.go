package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/userauth"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailLogic {
	return &GetUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserDetailLogic) GetUserDetail() (resp *types.GetUserDetailResp, err error) {

	uc, ok := userauth.GetUserFromCtx(l.ctx)
	if !ok {
		err = errorx.ErrorTokenInvalid()
		return nil, err
	}

	userInfo := &model.UserInfo{}
	if err = l.svcCtx.DB.Model(&model.UserInfo{}).
		Where("username = ?", uc.Username).
		First(userInfo).
		Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resp = &types.GetUserDetailResp{}
	_ = copier.Copy(resp, userInfo)
	resp.CreatedAt = userInfo.CreatedAt.UnixMilli()
	resp.UpdatedAt = userInfo.UpdatedAt.UnixMilli()

	resp.Role = &types.Role{}

	role := &model.Role{}
	if err = l.svcCtx.DB.Model(&model.Role{}).
		Where("id = ?", userInfo.RoleID).
		First(role).
		Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	_ = copier.Copy(resp.Role, role)
	resp.Role.CreatedAt = role.CreatedAt.UnixMilli()
	resp.Role.UpdatedAt = role.UpdatedAt.UnixMilli()

	return
}
