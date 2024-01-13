// Code generated by goctl. DO NOT EDIT.
package types

type AuthLoginReq struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captcha_id"`
}

type AuthLoginResp struct {
	Token string `json:"token"`
}

type GetAuthCaptchaResp struct {
	CaptchaId string `json:"captcha_id"`
	B64s      string `json:"b64s"`
}

type AuthLogoutReq struct {
	Token string `json:"token"`
}

type AddUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AddUserResp struct {
	ID int64 `json:"id"`
}

type GetUserPageReq struct {
	Page           int    `form:"page,default=1"`
	Size           int    `form:"size,default=10"`
	Username       string `form:"username,optional"` // 学号
	Name           string `form:"name,optional"`     // 姓名
	Gender         int    `form:"gender,optional"`
	CfId           string `form:"cf_id,optional"`
	AtcId          string `form:"atc_id,optional"`
	EnrollmentYear string `form:"enrollment_year,optional"`
}

type User struct {
	ID             int64  `json:"id"`              // 序号
	CreatedAt      string `json:"created_at"`      // 创建时间
	UpdatedAt      string `json:"updated_at"`      // 更新时间
	Username       string `json:"username"`        // 账号
	Email          string `json:"email"`           // 邮箱
	Avatar         string `json:"avatar"`          // 头像
	Name           string `json:"name"`            // english name
	Cname          string `json:"cname"`           // chinese name
	CfID           string `json:"cf_id"`           // codeforces用户名
	CfRating       int32  `json:"cf_rating"`       // codeforces rating
	CfRank         string `json:"cf_rank"`         // codeforces rank
	AtcID          string `json:"atc_id"`          // atcoder用户名
	NowcoderID     string `json:"nowcoder_id"`     // 牛客网id
	Gender         int32  `json:"gender"`          // 性别:0-未知;1-男;2-女
	IsEnable       bool   `json:"is_enable"`       // 是否启用:0-禁用;1-启用
	EnrollmentYear int32  `json:"enrollment_year"` // 入学年份
}

type GetUserPageResp struct {
	Items []*User `json:"items"`
	Total int64   `json:"total"`
}

type AddResourceReq struct {
	Name      string `json:"name"`
	Code      string `json:"code"`
	Type      string `json:"type"`
	ParentID  int    `json:"parent_id"`
	Path      string `json:"path"`
	Icon      string `json:"icon,optional"`
	Component string `json:"component,optional"`
	IsShow    bool   `json:"is_show"`
	IsEnable  bool   `json:"is_enable"`
	Order     int    `json:"order"`
}

type AddResourceResp struct {
	ID int64 `json:"id"`
}

type SaveResourceReq struct {
	ID        int64  `path:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Type      string `json:"type"`
	ParentID  int    `json:"parent_id"`
	Path      string `json:"path"`
	Icon      string `json:"icon,optional"`
	Component string `json:"component,optional"`
	IsShow    bool   `json:"is_show"`
	IsEnable  bool   `json:"is_enable"`
	Order     int    `json:"order"`
}

type SaveResourceResp struct {
	ID int64 `json:"id"`
}

type DeleteResourceReq struct {
	ID int64 `path:"id"`
}

type DeleteResourceResp struct {
	DelectCount int `json:"delect_count"`
}

type GetResourceTreeResp struct {
	Resource []*Resource `json:"resource"`
}

type Resource struct {
	ID        int64       `json:"id"`
	Name      string      `json:"name"`
	Code      string      `json:"code"`
	Type      string      `json:"type"`
	ParentID  int         `json:"parent_id"`
	Path      string      `json:"path"`
	Icon      string      `json:"icon"`
	Component string      `json:"component"`
	IsShow    bool        `json:"is_show"`
	IsEnable  bool        `json:"is_enable"`
	Order     int         `json:"order"`
	Children  []*Resource `json:"children"`
}

type AddRoleReq struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	IsEnable bool   `json:"is_enable"`
}

type AddRoleResp struct {
	ID int64 `json:"id"`
}
