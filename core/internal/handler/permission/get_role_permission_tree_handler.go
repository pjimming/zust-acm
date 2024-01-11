package permission

import (
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/permission"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/utils/httpresp"
)

func GetRolePermissionTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := permission.NewGetRolePermissionTreeLogic(r.Context(), svcCtx)
		resp, err := l.GetRolePermissionTree()

		httpresp.HTTP(w, r, resp, err)

	}
}
