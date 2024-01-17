package competition

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/sqlbuilder"

	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompetitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCompetitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompetitionLogic {
	return &GetCompetitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCompetitionLogic) GetCompetition(req *types.GetCompetitionReq) (resp *types.GetCompetitionResp, err error) {

	resp = &types.GetCompetitionResp{
		Items: make([]*types.Competition, 0),
		Total: 0,
	}

	sb := sqlbuilder.NewSQLBuilder(l.svcCtx.DB.Model(&model.Competition{})).
		AndStringLike("name", req.Name).
		AndStringIn("type", req.Type).
		AndIntIn("season_year", req.SeasonYear).
		OrderDesc("updated_at").
		ToSession()

	competitions, count, err := dao.Competition.GetPage(sb, req.Page, req.Size)
	if err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resp.Total = count

	for _, competition := range competitions {
		tmp := &types.Competition{}
		_ = copier.Copy(tmp, competition)
		tmp.CreatedAt = competition.CreatedAt.UnixMilli()
		tmp.UpdatedAt = competition.UpdatedAt.UnixMilli()
		tmp.StartTime = competition.StartTime.UnixMilli()
		tmp.EndTime = competition.EndTime.UnixMilli()
		resp.Items = append(resp.Items, tmp)
	}

	return
}
