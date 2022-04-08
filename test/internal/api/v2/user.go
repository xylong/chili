package v2

import (
	"github.com/xylong/chili"
	"net/http"
)

// Logoff 注销
func Logoff(ctx *chili.Context) {
	ctx.String(http.StatusOK, "version 2")
}
