package httpreq

import (
	"github.com/go-playground/form"
	"net/http"
)

// ParseForm parses the form request.
func ParseForm(r *http.Request, v interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	if err := form.NewDecoder().Decode(v, r.Form); err != nil {
		return err
	}
	return nil
}
