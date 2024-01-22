package team

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/constant"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddTeamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTeamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTeamLogic {
	return &AddTeamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTeamLogic) AddTeam(req *types.AddTeamReq) error {

	team := &model.Team{}
	_ = copier.Copy(team, req)
	// todo: custom trans

	if err := dao.Team.Insert(l.svcCtx.DB, team); err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	userTeamRels := make([]*model.UserTeamRel, 0)
	userTeamRels = append(userTeamRels, &model.UserTeamRel{
		UserID: req.LeaderID,
		TeamID: team.ID,
		Type:   constant.TeamLeader,
	})
	for _, memberId := range req.MemberIds {
		userTeamRels = append(userTeamRels, &model.UserTeamRel{
			UserID: memberId,
			TeamID: team.ID,
			Type:   constant.TeamMember,
		})
	}
	for _, teacherId := range req.TeacherIds {
		userTeamRels = append(userTeamRels, &model.UserTeamRel{
			UserID: teacherId,
			TeamID: team.ID,
			Type:   constant.TeamTeacher,
		})
	}

	if err := dao.UserTeamRel.BatchAdd(l.svcCtx.DB, userTeamRels); err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	return nil
}
