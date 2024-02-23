syntax = "v1"

info(
	title: "{{convertToUpperCamelCase .Model}}"
	desc: "{{convertToUpperCamelCase .Model}} apis"
	author: "panjiangming"
	email: "panjm2001@126.com"
	version: "v1"
)

type (
	{{firstUpper .Model}} {
		ID        int64  `json:"id"`         // 序号
		CreatedAt int64  `json:"created_at"` // 创建时间
		UpdatedAt int64  `json:"updated_at"` // 更新时间
		{{range .Fields -}}
            {{if and (ne .Name "id") (ne .Name "created_at") (ne .Name "updated_at") (ne .Name "deleted_at")  (ne .Name "deleted")  (ne .Name "created")  (ne .Name "updated") -}}
                {{convertToCamelCase .Name}}	{{.Type}} {{.Tag}}
            {{else -}}
            {{end}}
        {{- end}}
	}

	Add{{firstUpper .Model}}Req {
		{{range .Fields -}}
            {{if and (ne .Name "id") (ne .Name "created_at") (ne .Name "updated_at") (ne .Name "deleted_at")  (ne .Name "deleted")  (ne .Name "created")  (ne .Name "Updated") -}}
                {{convertToCamelCase .Name}}	{{.Type}} {{.Tag}}
            {{else -}}
            {{end}}
        {{- end}}
	}

	Get{{firstUpper .Model}}Req {
		Page  int    `form:"page,default=1"`
		Size  int    `form:"size,default=10"`
		// todo: complete
	}

	Get{{firstUpper .Model}}Resp {
		Items []*{{firstUpper .Model}} `json:"items"`
		Total int64      `json:"total"`
	}

	Update{{firstUpper .Model}}Req {
		ID     int64  `path:"id"`              // 序号
		{{range .Fields -}}
            {{if and (ne .Name "id") (ne .Name "created_at") (ne .Name "updated_at") (ne .Name "deleted_at")  (ne .Name "deleted")  (ne .Name "created")  (ne .Name "Updated") -}}
                {{convertToCamelCase .Name}}	{{.Type}} {{.Tag}}
            {{else -}}
            {{end}}
        {{- end}}
	}

	Delete{{firstUpper .Model}}Req {
		ID int64 `path:"id"`
	}
)

@server(
	group: {{toLower .Model}}
	middleware: JwtAuth
)
service core-api {
	@doc "新增"
	@handler Add{{firstUpper .Model}}
	post /api/v1/{{toLower .Model}} (Add{{firstUpper .Model}}Req)

	@doc "分页查询"
	@handler Get{{firstUpper .Model}}
	get /api/v1/{{toLower .Model}} (Get{{firstUpper .Model}}Req) returns (Get{{firstUpper .Model}}Resp)

	@doc "更新信息"
	@handler Update{{firstUpper .Model}}
	put /api/v1/{{toLower .Model}}/:id (Update{{firstUpper .Model}}Req)

	@doc "删除"
	@handler Delete{{firstUpper .Model}}
	delete /api/v1/{{toLower .Model}}/:id (Delete{{firstUpper .Model}}Req)
}