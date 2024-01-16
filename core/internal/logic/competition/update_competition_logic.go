package competition

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"time"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCompetitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCompetitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCompetitionLogic {
	return &UpdateCompetitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCompetitionLogic) UpdateCompetition(req *types.UpdateCompetitionReq) error {

	competition, err := dao.Competition.FindOne(l.svcCtx.DB, req.ID)
	if err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	_ = copier.Copy(competition, req)
	competition.StartTime = time.UnixMilli(req.StartTime)
	competition.EndTime = time.UnixMilli(req.EndTime)

	if err = dao.Competition.UpdateOne(l.svcCtx.DB, competition); err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	return nil
}
