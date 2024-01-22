package team

import (
	"context"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/constant"

	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/utils/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTeamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTeamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTeamLogic {
	return &UpdateTeamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTeamLogic) UpdateTeam(req *types.UpdateTeamReq) error {

	team, err := dao.Team.FindOne(l.svcCtx.DB, req.ID)
	if err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	_ = copier.Copy(team, req)

	if err = dao.Team.UpdateOne(l.svcCtx.DB, team); err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	if err = dao.UserTeamRel.RemoveByTeamIdsForce(l.svcCtx.DB, []int64{req.ID}); err != nil {
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

	if err = dao.UserTeamRel.BatchAdd(l.svcCtx.DB, userTeamRels); err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	return nil
}
