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

	if isExist, _ := l.svcCtx.UserAuthModel.FindOneByUsername(l.ctx, req.Username); isExist != nil {
		err = errorx.Error400f("[%s]用户已经存在", req.Username)
		return nil, err
	}

	user := &model.UserAuth{
		Username: req.Username,
	}
	user.Password, _ = userauth.EncryptPwd(req.Password)
	if _, err = l.svcCtx.UserAuthModel.Insert(l.ctx, user); err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resp = &types.AddUserResp{ID: int64(user.Id)}

	return
}
