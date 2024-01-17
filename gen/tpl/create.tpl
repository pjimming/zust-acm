package {{toLower .Model}}

import (
	"context"

	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/model"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type Add{{firstUpper .Model}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdd{{firstUpper .Model}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Add{{firstUpper .Model}}Logic {
	return &Add{{firstUpper .Model}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Add{{firstUpper .Model}}Logic) Add{{firstUpper .Model}}(req *types.Add{{firstUpper .Model}}Req) error {

	{{.Model}} := &model.{{firstUpper .Model}}{}
	_ = copier.Copy({{.Model}}, req)
	// todo: custom trans

	if err := dao.{{firstUpper .Model}}.Insert(l.svcCtx.DB, {{.Model}}); err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	return nil
}
