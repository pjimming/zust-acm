package {{toLower .Model}}

import (
	"context"

	"github.com/pjimming/zustacm/core/dao"
	"github.com/pjimming/zustacm/core/utils/errorx"
	"github.com/pjimming/zustacm/core/internal/svc"
	"github.com/pjimming/zustacm/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Delete{{firstUpper .Model}}Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelete{{firstUpper .Model}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Delete{{firstUpper .Model}}Logic {
	return &Delete{{firstUpper .Model}}Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Delete{{firstUpper .Model}}Logic) Delete{{firstUpper .Model}}(req *types.Delete{{firstUpper .Model}}Req) error {

	err := dao.{{firstUpper .Model}}.DeleteOne(l.svcCtx.DB, req.ID)
	if err != nil {
		err = errorx.ErrorDB(err)
		return err
	}

	return nil
}
