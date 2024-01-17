package record

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

type GetRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecordLogic {
	return &GetRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRecordLogic) GetRecord(req *types.GetRecordReq) (resp *types.GetRecordResp, err error) {

	resp = &types.GetRecordResp{
		Items: make([]*types.Record, 0),
		Total: 0,
	}

	sb := sqlbuilder.NewSQLBuilder(l.svcCtx.DB.Model(&model.Competition{})).
		AndStringIn("type", req.Type).
		AndIntIn("competition_id", req.CompetitionID).
		OrderDesc("updated_at").
		ToSession()

	records, count, err := dao.Record.GetPage(sb, req.Page, req.Size)
	if err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resp.Total = count

	competitionIds := make([]int64, 0)
	for _, record := range records {
		competitionIds = append(competitionIds, record.CompetitionID)
	}

	competitions, err := dao.Competition.FindByIds(l.svcCtx.DB, competitionIds)
	if err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	competitionMap := make(map[int64]*model.Competition)
	for _, competition := range competitions {
		competitionMap[competition.ID] = competition
	}

	for _, record := range records {
		tmp := &types.Record{}
		_ = copier.Copy(tmp, record)
		tmp.CompetitionName = competitionMap[record.CompetitionID].Name
		tmp.CompetitionType = competitionMap[record.CompetitionID].Type
		tmp.SeasonYear = competitionMap[record.CompetitionID].SeasonYear
		tmp.CreatedAt = record.CreatedAt.UnixMilli()
		tmp.UpdatedAt = record.UpdatedAt.UnixMilli()
		resp.Items = append(resp.Items, tmp)
	}

	return
}
