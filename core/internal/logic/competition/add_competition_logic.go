package competition

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"time"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCompetitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCompetitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCompetitionLogic {
	return &AddCompetitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCompetitionLogic) AddCompetition(req *types.AddCompetitionReq) (resp *types.AddCompetitionResp, err error) {

	competition := &model.Competition{}
	_ = copier.Copy(competition, req)
	competition.StartTime = time.UnixMilli(req.StartTime)
	competition.EndTime = time.UnixMilli(req.EndTime)

	if err = dao.Competition.Insert(l.svcCtx.DB, competition); err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resp = &types.AddCompetitionResp{ID: competition.ID}

	return
}
