package {{toLower .Model}}

import (
	"context"

	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/utils/sqlbuilder"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Get{{firstUpper .Model}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGet{{firstUpper .Model}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Get{{firstUpper .Model}}Logic {
	return &Get{{firstUpper .Model}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get{{firstUpper .Model}}Logic) Get{{firstUpper .Model}}(req *types.Get{{firstUpper .Model}}Req) (resp *types.Get{{firstUpper .Model}}Resp, err error) {

	resp = &types.Get{{firstUpper .Model}}Resp{
		Items: make([]*types.{{firstUpper .Model}}, 0),
		Total: 0,
	}

	sb := sqlbuilder.NewSQLBuilder(l.svcCtx.DB.Model(&model.{{firstUpper .Model}}{})).
		// todo: custom query
		OrderDesc("updated_at").
		ToSession()

	{{.Model}}s, count, err := dao.{{firstUpper .Model}}.GetPage(sb, req.Page, req.Size)
	if err != nil {
		err = errorx.ErrorDB(err)
		return nil, err
	}

	resp.Total = count

	for _, {{.Model}} := range {{.Model}}s {
		resp.Items = append(resp.Items, modelToTypes({{.Model}}))
	}

	return
}
