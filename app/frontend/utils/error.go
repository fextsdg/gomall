package utils

import "github.com/cloudwego/hertz/pkg/common/hlog"

func MustHandlerError(err error) {
	hlog.Fatal(err)
}
