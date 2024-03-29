package auth

import (
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/httpresp"
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/auth"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AuthLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpresp.HttpError(w, r, errorx.NewCodeError(2, http.StatusBadRequest, err.Error()))
			return
		}

		l := auth.NewAuthLoginLogic(r.Context(), svcCtx)
		resp, err := l.AuthLogin(&req)

		httpresp.HTTP(w, r, resp, err)

	}
}
