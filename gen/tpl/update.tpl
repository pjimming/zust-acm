package {{toLower .Model}}

import (
	"context"

	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type Update{{firstUpper .Model}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdate{{firstUpper .Model}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Update{{firstUpper .Model}}Logic {
	return &Update{{firstUpper .Model}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Update{{firstUpper .Model}}Logic) Update{{firstUpper .Model}}(req *types.Update{{firstUpper .Model}}Req) error {

	{{.Model}}, err := dao.{{firstUpper .Model}}.FindOne(l.svcCtx.DB, req.ID)
	if err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	_ = copier.Copy({{.Model}}, req)
	// todo: custom trans

	if err = dao.{{firstUpper .Model}}.UpdateOne(l.svcCtx.DB, {{.Model}}); err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	return nil
}
