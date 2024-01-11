package resource

import (
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/resource"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/utils/httpresp"
)

func GetResourceTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := resource.NewGetResourceTreeLogic(r.Context(), svcCtx)
		resp, err := l.GetResourceTree()

		httpresp.HTTP(w, r, resp, err)

	}
}
