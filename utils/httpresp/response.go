package httpresp

import (
	"net/http"

	"github.com/pjimming/zustacm/utils/errorx"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	codeServerError     = 1
	internalServerError = "内部错误，请查看日志信息或联系开发者"
)

type httpResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HTTP(w http.ResponseWriter, r *http.Request, resp interface{}, err error) {
	if err != nil {
		HttpError(w, r, err)
	} else {
		HttpOkJson(w, r, resp)
	}
}

func HttpError(w http.ResponseWriter, r *http.Request, err error) {
	codeErr, ok := errorx.FromError(err)
	if ok {
		httpx.WriteJson(w, codeErr.Status(), httpResp{
			Code:    codeErr.Code(),
			Message: codeErr.Error(),
		})
	} else {
		httpx.WriteJson(w, http.StatusInternalServerError, httpResp{
			Code:    codeServerError,
			Message: internalServerError,
		})
		logx.WithContext(r.Context()).Error(err)
	}
}

func HttpOkJson(w http.ResponseWriter, r *http.Request, resp interface{}) {
	httpx.OkJsonCtx(r.Context(), w, httpResp{Code: http.StatusOK, Data: resp})
}
