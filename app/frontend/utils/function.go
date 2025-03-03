package utils

import "golang.org/x/net/context"

func GetUserIdFromCtx(ctx context.Context) int32 {
	userId := ctx.Value(SessionUserId)
	if userId == "" {
		return 0
	}
	return userId.(int32)
}
