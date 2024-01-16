package competition

import (
	"github.com/pjimming/zustacm/core/utils/httpreq"
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/competition"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/httpresp"
)

func GetCompetitionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCompetitionReq
		if err := httpreq.ParseForm(r, &req); err != nil {
			httpresp.HttpError(w, r, errorx.NewCodeError(2, http.StatusBadRequest, err.Error()))
			return
		}

		l := competition.NewGetCompetitionLogic(r.Context(), svcCtx)
		resp, err := l.GetCompetition(&req)

		httpresp.HTTP(w, r, resp, err)

	}
}
