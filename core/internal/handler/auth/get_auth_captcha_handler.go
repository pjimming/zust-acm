package auth

import (
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/auth"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/utils/httpresp"
)

func GetAuthCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewGetAuthCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetAuthCaptcha()

		httpresp.HTTP(w, r, resp, err)

	}
}
