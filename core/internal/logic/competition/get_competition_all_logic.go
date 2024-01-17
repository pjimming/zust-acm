package competition

import (
	"context"
	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/utils/errorx"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompetitionAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCompetitionAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompetitionAllLogic {
	return &GetCompetitionAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCompetitionAllLogic) GetCompetitionAll() (resp *types.GetCompetitionResp, err error) {

	resp = &types.GetCompetitionResp{Items: make([]*types.Competition, 0)}

	competitions, err := dao.Competition.FindAll(l.svcCtx.DB)
	if err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	for _, competition := range competitions {
		resp.Items = append(resp.Items, modelToTypes(competition))
	}

	return
}
