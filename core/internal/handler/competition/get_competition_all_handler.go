package competition

import (
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/competition"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/utils/httpresp"
)

func GetCompetitionAllHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := competition.NewGetCompetitionAllLogic(r.Context(), svcCtx)
		resp, err := l.GetCompetitionAll()

		httpresp.HTTP(w, r, resp.Items, err)

	}
}
