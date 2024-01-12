package helper

import (
	"time"

	"github.com/pjimming/zustacm/core/utils/constant"
)

func GetLocalDateTime(t time.Time) string {
	return t.Local().Format(constant.DateTimeLayout)
}
