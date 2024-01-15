package httpreq

import (
	"encoding/json"
	"net/url"
	"testing"
)

type GetUserPageReq struct {
	Page           int     `form:"page,default=1"`
	Size           int     `form:"size,default=10"`
	Username       string  `form:"username,optional"` // 学号
	Name           string  `form:"name,optional"`     // 姓名
	Gender         int     `form:"gender,optional"`
	CfId           string  `form:"cf_id,optional"`
	AtcId          string  `form:"atc_id,optional"`
	EnrollmentYear []int64 `form:"enrollment_year[],optional"`
}

func TestParseForm(t *testing.T) {

	// 解析 URL
	parsedURL, err := url.Parse("api/user?enrollment_year[]=2020&enrollment_year[]=2021&page=1&size=10")
	if err != nil {
		t.Fatal(err)
	}

	data := &GetUserPageReq{}

	// 获取查询参数
	t.Log(parsedURL.Query())

	json.Unmarshal([]byte(parsedURL.Query().Encode()), data)
	t.Log(data)
}
