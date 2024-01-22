package team

import (
	"github.com/pjimming/zustacm/core/utils/httpreq"
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/team"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/httpresp"
)

func GetTeamHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTeamReq
		if err := httpreq.ParseForm(r, &req); err != nil {
			httpresp.HttpError(w, r, errorx.NewCodeError(2, http.StatusBadRequest, err.Error()))
			return
		}

		l := team.NewGetTeamLogic(r.Context(), svcCtx)
		resp, err := l.GetTeam(&req)

		httpresp.HTTP(w, r, resp, err)

	}
}
