package user

import (
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/user"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/utils/httpresp"
)

func GetUserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetUserDetail()

		httpresp.HTTP(w, r, resp, err)

	}
}
