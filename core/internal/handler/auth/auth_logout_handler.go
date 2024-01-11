package auth

import (
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/auth"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/utils/httpresp"
)

func AuthLogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewAuthLogoutLogic(r.Context(), svcCtx)
		err := l.AuthLogout()

		httpresp.HTTP(w, r, nil, err)

	}
}
