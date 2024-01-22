package user

import (
	"context"
	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/utils/errorx"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAllLogic {
	return &GetUserAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserAllLogic) GetUserAll() (resp *types.GetUserPageResp, err error) {

	resp = &types.GetUserPageResp{Items: make([]*types.User, 0)}

	users, err := dao.UserInfo.FindAll(l.svcCtx.DB)
	if err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	for _, user := range users {
		resp.Items = append(resp.Items, model2Type(user))
	}

	return
}
