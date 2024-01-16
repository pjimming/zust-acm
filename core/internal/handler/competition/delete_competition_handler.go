package competition

import (
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/competition"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/httpresp"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteCompetitionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCompetitionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpresp.HttpError(w, r, errorx.NewCodeError(2, http.StatusBadRequest, err.Error()))
			return
		}

		l := competition.NewDeleteCompetitionLogic(r.Context(), svcCtx)
		err := l.DeleteCompetition(&req)

		httpresp.HTTP(w, r, nil, err)

	}
}
