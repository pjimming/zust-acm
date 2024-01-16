package codeforces

import (
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/codeforces"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/utils/httpresp"
)

func SyncCfAllHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := codeforces.NewSyncCfAllLogic(r.Context(), svcCtx)
		err := l.SyncCfAll()

		httpresp.HTTP(w, r, nil, err)

	}
}
