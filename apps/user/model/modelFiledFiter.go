package model

import (
	"data_source_management_center/common/tools"
	"fmt"
	"reflect"
	"strings"

	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	baseRemoveFiled = []string{"`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"}
)

// FiledOptional 用户字段表过滤
type FiledOptional struct {
	userAllFiledNames []string
	tableName         string
	tagFiledMap       map[string]string
	baseRemoveFiled   []string
}

func NewFiledOptional() *FiledOptional {
	return &FiledOptional{
		userAllFiledNames: userFieldNames,
		tableName:         "user",
		tagFiledMap:       make(map[string]string),
		baseRemoveFiled:   []string{"`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"},
	}
}

func (f *FiledOptional) getFiledTagName() {
	pType := reflect.TypeOf(&User{})
	fieldNum := pType.Elem().NumField()

	for i := 0; i < fieldNum; i++ {
		ftype := pType.Elem().Field(i)
		f.tagFiledMap[ftype.Tag.Get("db")] = ftype.Name
	}
}

func (f *FiledOptional) GenUpdate(removeFiled []string, user *User) (formatSql string, data []interface{}, err error) {
	//更新只能更新密码之外的信息
	f.baseRemoveFiled = append(f.baseRemoveFiled, "`password`")
	if len(removeFiled) != 0 {
		f.baseRemoveFiled = append(f.baseRemoveFiled, removeFiled...)
	}

	filterFiledRes := stringx.Remove(f.userAllFiledNames, f.baseRemoveFiled...)
	userFiled := strings.Join(filterFiledRes, ",")
	querySqlFormat := fmt.Sprintf("update %s (%s) values(", f.tableName, userFiled)

	var querySql strings.Builder
	querySql.WriteString(querySqlFormat)
	for i := 0; i < len(filterFiledRes); i++ {
		if i == (len(filterFiledRes) - 1) {
			querySql.WriteString("?)")
		} else {
			querySql.WriteString("?,")
		}
	}

	userValues := make([]interface{}, len(filterFiledRes))
	for _, v := range filterFiledRes {
		uValue, err := tools.GetFieldValueByTag("db", v, user)
		if err != nil {
			fmt.Println("获取值出错了:", err)
			continue
		}
		userValues = append(userValues, uValue.String())
	}

	return querySql.String(), userValues, nil

}

func RemoveFiled(fileds []string) string {
	baseRemoveFiled = append(baseRemoveFiled, fileds...)
	filterRes := strings.Join(stringx.Remove(userFieldNames, baseRemoveFiled...), "=?,") + "=?"
	return filterRes
}
