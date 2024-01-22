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

type AddResourceReq struct {
	Name      string `json:"name"`
	Code      string `json:"code"`
	Type      string `json:"type,optional"`
	ParentID  int    `json:"parent_id,optional"`
	Path      string `json:"path,optional"`
	Icon      string `json:"icon,optional"`
	Component string `json:"component,optional"`
	IsShow    bool   `json:"is_show,optional"`
	IsEnable  bool   `json:"is_enable,optional"`
	Order     int    `json:"order"`
}

type AddResourceResp struct {
	ID int64 `json:"id"`
}

type SaveResourceReq struct {
	ID        int64  `path:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Type      string `json:"type,optional"`
	ParentID  int    `json:"parent_id,optional"`
	Path      string `json:"path,optional"`
	Icon      string `json:"icon,optional"`
	Component string `json:"component,optional"`
	IsShow    bool   `json:"is_show,optional"`
	IsEnable  bool   `json:"is_enable,optional"`
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

type AddUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoldId   int64  `json:"role_id"`
	IsEnable bool   `json:"is_enable"`
}

type AddUserResp struct {
	ID int64 `json:"id"`
}

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

type User struct {
	ID             int64  `json:"id"`              // 序号
	CreatedAt      int64  `json:"created_at"`      // 创建时间
	UpdatedAt      int64  `json:"updated_at"`      // 更新时间
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
	Role           *Role  `json:"role"`            // 角色
}

type UserRole struct {
	ID       int64  `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	IsEnable bool   `json:"is_enable"`
}

type GetUserDetailResp struct {
	ID             int64  `json:"id"`              // 序号
	CreatedAt      int64  `json:"created_at"`      // 创建时间
	UpdatedAt      int64  `json:"updated_at"`      // 更新时间
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
	Role           *Role  `json:"role"`            // 角色
}

type GetUserPageResp struct {
	Items []*User `json:"items"`
	Total int64   `json:"total"`
}

type AssignRolesReq struct {
	ID       int64  `json:"id"`
	RoleId   int64  `json:"role_id"`
	Username string `json:"username"`
}

type AddRoleReq struct {
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	IsEnable    bool    `json:"is_enable"`
	ResourceIds []int64 `json:"resource_ids"`
}

type AddRoleResp struct {
	ID int64 `json:"id"`
}

type Role struct {
	ID          int64   `json:"id"`           // 序号
	CreatedAt   int64   `json:"created_at"`   // 创建时间
	UpdatedAt   int64   `json:"updated_at"`   // 更新时间
	Code        string  `json:"code"`         // 角色编码
	Name        string  `json:"name"`         // 角色名
	IsEnable    bool    `json:"is_enable"`    // 启用状态:0-禁用;1-启用
	ResourceIds []int64 `json:"resource_ids"` // 权限
}

type GetRoleAllReq struct {
	Code     string `form:"code,optional"`
	Name     string `form:"name,optional"`
	IsEnable bool   `form:"is_enable,default=false"`
}

type GetRoleAllResp struct {
	Items []*Role `json:"items"`
}

type GetRolePageReq struct {
	Page int    `form:"page,default=1"`
	Size int    `form:"size,default=10"`
	Code string `form:"code,optional"`
	Name string `form:"name,optional"`
}

type GetRolePageResp struct {
	Items []*Role `json:"items"`
	Total int64   `json:"total"`
}

type UpdateRoleReq struct {
	ID          int64   `path:"id"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	IsEnable    bool    `json:"is_enable"`
	ResourceIds []int64 `json:"resource_ids"`
}

type SyncUserCfReq struct {
	ID int64 `path:"id"`
}

type Competition struct {
	ID              int64  `json:"id"`               // 序号
	CreatedAt       int64  `json:"created_at"`       // 创建时间
	UpdatedAt       int64  `json:"updated_at"`       // 更新时间
	Name            string `json:"name"`             // 比赛名称
	Type            string `json:"type"`             // 比赛类型
	SeasonYear      int32  `json:"season_year"`      // 赛季年
	StartTime       int64  `json:"start_time"`       // 开始时间
	EndTime         int64  `json:"end_time"`         // 结束时间
	OfficialWebsite string `json:"official_website"` // 比赛官网
}

type AddCompetitionReq struct {
	Name            string `json:"name"`                      // 比赛名称
	Type            string `json:"type"`                      // 比赛类型
	SeasonYear      int    `json:"season_year"`               // 赛季年
	StartTime       int64  `json:"start_time"`                // 开始时间
	EndTime         int64  `json:"end_time"`                  // 结束时间
	OfficialWebsite string `json:"official_website,optional"` // 比赛官网
}

type AddCompetitionResp struct {
	ID int64 `json:"id"`
}

type GetCompetitionReq struct {
	Page       int      `form:"page,default=1"`
	Size       int      `form:"size,default=10"`
	Name       string   `form:"name,optional"`
	Type       []string `form:"type[],optional"`
	SeasonYear []int64  `form:"season_year[],optional"`
}

type GetCompetitionResp struct {
	Items []*Competition `json:"items"`
	Total int64          `json:"total"`
}

type UpdateCompetitionReq struct {
	ID              int64  `path:"id"`                        // 序号
	Name            string `json:"name"`                      // 比赛名称
	Type            string `json:"type"`                      // 比赛类型
	SeasonYear      int32  `json:"season_year"`               // 赛季年
	StartTime       int64  `json:"start_time"`                // 开始时间
	EndTime         int64  `json:"end_time"`                  // 结束时间
	OfficialWebsite string `json:"official_website,optional"` // 比赛官网
}

type DeleteCompetitionReq struct {
	ID int64 `path:"id"`
}

type Record struct {
	ID              int64  `json:"id"`               // 序号
	CreatedAt       int64  `json:"created_at"`       // 创建时间
	UpdatedAt       int64  `json:"updated_at"`       // 更新时间
	CompetitionID   int64  `json:"competition_id"`   // 竞赛id
	CompetitionName string `json:"competition_name"` // 竞赛全称
	CompetitionType string `json:"competition_type"` // 竞赛类型
	SeasonYear      int32  `json:"season_year"`      // 赛季年
	Type            string `json:"type"`             // 奖项类型
	Remark          string `json:"remark"`           // 备注
}

type AddRecordReq struct {
	CompetitionID int64  `json:"competition_id"`  // 竞赛id
	Type          string `json:"type"`            // 奖项类型
	Remark        string `json:"remark,optional"` // 备注
}

type GetRecordReq struct {
	Page          int      `form:"page,default=1"`
	Size          int      `form:"size,default=10"`
	CompetitionID []int64  `form:"competition_id[],optional"` // 竞赛id
	Type          []string `form:"type[],optional"`           // 奖项类型
}

type GetRecordResp struct {
	Items []*Record `json:"items"`
	Total int64     `json:"total"`
}

type UpdateRecordReq struct {
	ID            int64  `path:"id"`              // 序号
	CompetitionID int64  `json:"competition_id"`  // 竞赛id
	Type          string `json:"type"`            // 奖项类型
	Remark        string `json:"remark,optional"` // 备注
}

type DeleteRecordReq struct {
	ID int64 `path:"id"`
}

type SysDict struct {
	ID        int64  `json:"id"`         // 序号
	CreatedAt int64  `json:"created_at"` // 创建时间
	UpdatedAt int64  `json:"updated_at"` // 更新时间
	Label     string `json:"label"`      // 标签
	Value     string `json:"value"`      // 值
	Type      string `json:"type"`       // 类型
	Remark    string `json:"remark"`     // 备注
}

type AddSysDictReq struct {
	Label  string `json:"label"`           // 标签
	Value  string `json:"value"`           // 值
	Type   string `json:"type"`            // 类型
	Remark string `json:"remark,optional"` // 备注
}

type GetSysDictReq struct {
	Page  int    `form:"page,default=1"`
	Size  int    `form:"size,default=10"`
	Label string `form:"label,optional"`
	Value string `form:"value,optional"`
	Type  string `form:"type,optional"` // 字典类型
}

type GetSysDictResp struct {
	Items []*SysDict `json:"items"`
	Total int64      `json:"total"`
}

type UpdateSysDictReq struct {
	ID     int64  `path:"id"`              // 序号
	Label  string `json:"label"`           // 标签
	Value  string `json:"value"`           // 值
	Type   string `json:"type"`            // 类型
	Remark string `json:"remark,optional"` // 备注
}

type DeleteSysDictReq struct {
	ID int64 `path:"id"`
}

type GetSysDictTypesResp struct {
	Types []string `json:"types"`
}

type Team struct {
	ID        int64         `json:"id"`         // 序号
	CreatedAt int64         `json:"created_at"` // 创建时间
	UpdatedAt int64         `json:"updated_at"` // 更新时间
	Leader    *TeamMember   `json:"leader"`     // 队长
	Members   []*TeamMember `json:"members"`    // 队员
	Teachers  []*TeamMember `json:"teachers"`   // 指导教师
	Name      string        `json:"name"`       // 队伍名称
	Remark    string        `json:"remark"`     // 备注
}

type TeamMember struct {
	ID       int64  `json:"id"` // user_info.id
	Username string `json:"username"`
	Cname    string `json:"cname"`
}

type AddTeamReq struct {
	LeaderID   int64   `json:"leader_id"`            // 队长id；user_info.id
	Name       string  `json:"name"`                 // 队伍名称
	MemberIds  []int64 `json:"member_ids,optional"`  // 成员
	TeacherIds []int64 `json:"teacher_ids,optional"` // 指导老师
	Remark     string  `json:"remark,optional"`      // 备注
}

type GetTeamReq struct {
	Page      int     `form:"page,default=1"`
	Size      int     `form:"size,default=10"`
	Name      string  `form:"name,optional"` // 队伍名称
	MemberIds []int64 `form:"member_ids[],optional"`
}

type GetTeamResp struct {
	Items []*Team `json:"items"`
	Total int64   `json:"total"`
}

type UpdateTeamReq struct {
	ID         int64   `path:"id"`                   // 序号
	LeaderID   int64   `json:"leader_id"`            // 队长id；user_info.id
	Name       string  `json:"name"`                 // 队伍名称
	MemberIds  []int64 `json:"member_ids,optional"`  // 成员
	TeacherIds []int64 `json:"teacher_ids,optional"` // 指导老师
	Remark     string  `json:"remark,optional"`      // 备注
}

type DeleteTeamReq struct {
	ID int64 `path:"id"`
}
