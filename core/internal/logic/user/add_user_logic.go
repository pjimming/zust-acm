package user

import (
	"context"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/userauth"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserReq) (resp *types.AddUserResp, err error) {

	var userCount int64
	if err = l.svcCtx.DB.Model(&model.UserAuth{}).
		Where("username = ?", req.Username).
		Count(&userCount).Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	if userCount > 0 {
		err = errorx.Error400f("[%s]用户已经存在", req.Username)
		return nil, err
	}

	userAuth := &model.UserAuth{Username: req.Username}
	userInfo := &model.UserInfo{Username: req.Username, RoleID: req.RoldId, IsEnable: req.IsEnable}
	userAuth.Password, _ = userauth.EncryptPwd(req.Password)

	if err = l.svcCtx.DB.Create(userAuth).Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}
	if err = l.svcCtx.DB.Create(userInfo).Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resp = &types.AddUserResp{ID: userAuth.ID}

	return
}
