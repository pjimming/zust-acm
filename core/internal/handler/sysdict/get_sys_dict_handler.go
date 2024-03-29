package sysdict

import (
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/sysdict"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/httpresp"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSysDictHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSysDictReq
		if err := httpx.Parse(r, &req); err != nil {
			httpresp.HttpError(w, r, errorx.NewCodeError(2, http.StatusBadRequest, err.Error()))
			return
		}

		l := sysdict.NewGetSysDictLogic(r.Context(), svcCtx)
		resp, err := l.GetSysDict(&req)

		httpresp.HTTP(w, r, resp, err)

	}
}
