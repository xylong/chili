package v1

import (
	"github.com/xylong/chili"
	"net/http"
)

// Store 保存文章
func Store(ctx *chili.Context) {
	ctx.JSON(http.StatusOK, "创建")
}
