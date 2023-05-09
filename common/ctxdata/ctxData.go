package ctxdata

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		fmt.Println("获取的用户id：", jsonUid)
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	}
	return uid
}
