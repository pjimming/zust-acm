package cfhelper

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

var cfHelper = NewCfHelper(resty.New())

func TestCfHelper_GetUserInfo(t *testing.T) {
	ast := assert.New(t)
	resp, err := cfHelper.GetUserInfo([]string{"Payxtl", "tourist"})
	ast.Nil(err)
	ast.NotNil(resp)
}
