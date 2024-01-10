package auth

import (
	"github.com/pjimming/zustacm/core/utils/httpresp"
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/auth"
	"github.com/pjimming/zustacm/core/internal/svc"
)

func GetAuthCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewGetAuthCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetAuthCaptcha()

		httpresp.HTTP(w, r, resp, err)

	}
}
