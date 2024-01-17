package competition

import (
	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/model"

	"github.com/jinzhu/copier"
)

func modelToTypes(from *model.Competition) *types.Competition {
	tmp := &types.Competition{}
	_ = copier.Copy(tmp, from)
	tmp.CreatedAt = from.CreatedAt.UnixMilli()
	tmp.UpdatedAt = from.UpdatedAt.UnixMilli()
	tmp.StartTime = from.StartTime.UnixMilli()
	tmp.EndTime = from.EndTime.UnixMilli()
	return tmp
}
