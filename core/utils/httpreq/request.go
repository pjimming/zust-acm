package httpreq

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/form"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// ParseForm parses the form request.
func ParseForm(r *http.Request, v interface{}) error {
	if err := httpx.ParsePath(r, v); err != nil {
		return err
	}
	if err := httpx.ParseHeaders(r, v); err != nil {
		return err
	}
	if err := httpx.ParseJsonBody(r, v); err != nil {
		return err
	}

	if err := r.ParseForm(); err != nil {
		return err
	}
	if err := form.NewDecoder().Decode(v, r.Form); err != nil {
		return err
	}

	return nil
}

func ParseJson(r *http.Request, v interface{}) error {
	if err := httpx.ParseHeaders(r, v); err != nil {
		return err
	}
	if err := httpx.ParsePath(r, v); err != nil {
		return err
	}
	if err := httpx.ParseForm(r, v); err != nil {
		return err
	}

	// get RequestBody
	reqBody, _ := io.ReadAll(r.Body)
	_ = r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(reqBody))

	return json.Unmarshal(reqBody, v)
}
