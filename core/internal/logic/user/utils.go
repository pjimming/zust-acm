package user

import (
	"github.com/jinzhu/copier"
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/model"
)

func model2Type(from *model.UserInfo) *types.User {
	tmp := &types.User{}
	_ = copier.Copy(tmp, from)
	tmp.CreatedAt = from.CreatedAt.UnixMilli()
	tmp.UpdatedAt = from.UpdatedAt.UnixMilli()
	// todo: custom trans
	return tmp
}
