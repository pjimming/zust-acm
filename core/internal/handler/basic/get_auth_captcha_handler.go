package basic

import (
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/basic"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/utils/httpresp"
)

func GetAuthCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := basic.NewGetAuthCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetAuthCaptcha()

		httpresp.HTTP(w, r, resp, err)

	}
}
