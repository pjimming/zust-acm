package codeforces

import (
	"context"
	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/utils/errorx"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncUserCfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncUserCfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncUserCfLogic {
	return &SyncUserCfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncUserCfLogic) SyncUserCf(req *types.SyncUserCfReq) error {

	userInfo, err := dao.UserInfo.FindOne(l.svcCtx.DB, req.ID)
	if err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	cfResp, err := l.svcCtx.CfHelper.GetUserInfo(userInfo.CfID)
	if err != nil {
		err = errorx.Error500f("do CF API fail, %v", err)
		return err
	}

	userInfo.CfRating = int32(cfResp.Rating)
	userInfo.CfRank = cfResp.Rank

	if err = dao.UserInfo.UpdateOne(l.svcCtx.DB, userInfo); err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	return nil
}
