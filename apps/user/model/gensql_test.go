package model

import (
	"fmt"
	"testing"
)

func TestGensql(t *testing.T) {
	u := User{
		Id:       11,
		Username: "测试名",
		Password: "1233434",
		Sex:      "男",
		Email:    "132234@qq.com",
		Info:     "ffsdfdfdsfsdfsdfd",
	}
	removeList := []string{"`sex`"}
	a := NewFiledOptional()
	updateSql, v, err := a.GenUpdate(removeList, &u)
	if err != nil {
		return
	}
	fmt.Println(updateSql)
	fmt.Println(v)

}
