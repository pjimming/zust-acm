package resource

import (
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/resource"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/utils/httpresp"
)

func GetResourceMenuTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := resource.NewGetResourceMenuTreeLogic(r.Context(), svcCtx)
		resp, err := l.GetResourceMenuTree()

		httpresp.HTTP(w, r, resp, err)

	}
}
