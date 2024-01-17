package {{toLower .Model}}

import (
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/model"

	"github.com/jinzhu/copier"
)

func modelToTypes(from *model.{{firstUpper .Model}}) *types.{{firstUpper .Model}} {
	tmp := &types.{{firstUpper .Model}}{}
	_ = copier.Copy(tmp, from)
	tmp.CreatedAt = from.CreatedAt.UnixMilli()
    tmp.UpdatedAt = from.UpdatedAt.UnixMilli()
	// todo: custom trans
	return tmp
}
