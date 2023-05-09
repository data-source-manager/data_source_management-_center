package model

import (
	"strings"

	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	baseRemoveFiled = []string{"`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"}
)

func RemoveFiled(fileds []string) string {
	baseRemoveFiled = append(baseRemoveFiled, fileds...)
	filterRes := strings.Join(stringx.Remove(userFieldNames, baseRemoveFiled...), "=?,") + "=?"
	return filterRes
}
