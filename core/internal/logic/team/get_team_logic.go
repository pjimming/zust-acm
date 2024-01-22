package team

import (
	"context"
	"github.com/pjimming/zustacm/core/utils/constant"

	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/sqlbuilder"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTeamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeamLogic {
	return &GetTeamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTeamLogic) GetTeam(req *types.GetTeamReq) (resp *types.GetTeamResp, err error) {

	resp = &types.GetTeamResp{
		Items: make([]*types.Team, 0),
		Total: 0,
	}

	uTRs := make([]*model.UserTeamRel, 0)
	if err = sqlbuilder.NewSQLBuilder(l.svcCtx.DB.Model(&model.UserTeamRel{})).
		AndIntIn("user_id", req.MemberIds).
		ToSession().
		Distinct("team_id").
		Find(&uTRs).
		Error; err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	teamIds := make([]int64, 0)
	for _, utr := range uTRs {
		teamIds = append(teamIds, utr.TeamID)
	}

	sb := sqlbuilder.NewSQLBuilder(l.svcCtx.DB.Model(&model.Team{})).
		AndStringLike("name", req.Name).
		AndIntIn("id", teamIds).
		OrderDesc("updated_at").
		ToSession()

	teams, count, err := dao.Team.GetPage(sb, req.Page, req.Size)
	if err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resp.Total = count

	for _, team := range teams {
		tmp := modelToTypes(team)
		utrs, err := dao.UserTeamRel.FindByTeamIds(l.svcCtx.DB, []int64{team.ID})
		if err != nil {
			err = errorx.ErrorDB(err)
			return nil, err
		}

		for _, utr := range utrs {
			userInfo, err := dao.UserInfo.FindOne(l.svcCtx.DB, utr.UserID)
			if err != nil {
				err = errorx.ErrorDB(err)
				return nil, err
			}

			switch utr.Type {
			case constant.TeamLeader:
				tmp.LeaderId = userInfo.ID
			case constant.TeamMember:
				tmp.MemberIds = append(tmp.MemberIds, userInfo.ID)
			case constant.TeamTeacher:
				tmp.TeacherIds = append(tmp.TeacherIds, userInfo.ID)
			}

		}

		resp.Items = append(resp.Items, tmp)
	}

	return
}
