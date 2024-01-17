package sysdict

import (
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/sysdict"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/utils/httpresp"
)

func GetSysDictTypesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := sysdict.NewGetSysDictTypesLogic(r.Context(), svcCtx)
		resp, err := l.GetSysDictTypes()

		httpresp.HTTP(w, r, resp.Types, err)

	}
}
