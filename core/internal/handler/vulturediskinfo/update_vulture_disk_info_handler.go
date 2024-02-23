package vulturediskinfo

import (
	"net/http"

	"github.com/pjimming/zustacm/core/internal/logic/vulturediskinfo"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/httpresp"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateVultureDiskInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateVultureDiskInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpresp.HttpError(w, r, errorx.NewCodeError(2, http.StatusBadRequest, err.Error()))
			return
		}

		l := vulturediskinfo.NewUpdateVultureDiskInfoLogic(r.Context(), svcCtx)
		err := l.UpdateVultureDiskInfo(&req)

		httpresp.HTTP(w, r, nil, err)

	}
}
