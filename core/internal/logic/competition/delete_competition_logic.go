package competition

import (
	"context"
	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/utils/errorx"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCompetitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCompetitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCompetitionLogic {
	return &DeleteCompetitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCompetitionLogic) DeleteCompetition(req *types.DeleteCompetitionReq) error {

	err := dao.Competition.DeleteOne(l.svcCtx.DB, req.ID)
	if err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	return nil
}
