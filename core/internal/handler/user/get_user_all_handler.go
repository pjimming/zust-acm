package user

import (
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/user"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/utils/httpresp"
)

func GetUserAllHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserAllLogic(r.Context(), svcCtx)
		resp, err := l.GetUserAll()

		httpresp.HTTP(w, r, resp.Items, err)

	}
}
