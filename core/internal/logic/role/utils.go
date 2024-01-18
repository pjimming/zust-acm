package role

import (
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/model"

	"github.com/jinzhu/copier"
)

func model2Types(from *model.Role) *types.Role {
	tmp := &types.Role{}
	_ = copier.Copy(tmp, from)
	tmp.CreatedAt = from.CreatedAt.UnixMilli()
	tmp.UpdatedAt = from.UpdatedAt.UnixMilli()
	// todo: custom trans
	return tmp
}
