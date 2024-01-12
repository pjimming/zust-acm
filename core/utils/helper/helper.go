package helper

import (
	"strconv"
	"strings"
	"time"

	"github.com/pjimming/zustacm/core/utils/constant"
)

func GetLocalDateTime(t time.Time) string {
	return t.Local().Format(constant.DateTimeLayout)
}

func ConvertStringToIntArray(s string) []int {
	split := strings.Split(s, ",")
	ret := make([]int, 0)
	for _, item := range split {
		num, err := strconv.Atoi(item)
		if err != nil {
			continue
		}
		ret = append(ret, num)
	}
	return ret
}

func ConvertStringToInt64Array(s string) []int64 {
	array := ConvertStringToIntArray(s)
	ret := make([]int64, 0)
	for _, item := range array {
		ret = append(ret, int64(item))
	}
	return ret
}
