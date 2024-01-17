package record

import (
	"github.com/pjimming/zustacm/core/utils/httpreq"
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/record"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/httpresp"
)

func GetRecordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRecordReq
		if err := httpreq.ParseForm(r, &req); err != nil {
			httpresp.HttpError(w, r, errorx.NewCodeError(2, http.StatusBadRequest, err.Error()))
			return
		}

		l := record.NewGetRecordLogic(r.Context(), svcCtx)
		resp, err := l.GetRecord(&req)

		httpresp.HTTP(w, r, resp, err)

	}
}
