package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xylong/chili/test/internal/request"
	"net/http"
)

func Login(ctx *gin.Context, param *request.LoginParam) {
	ctx.JSON(http.StatusOK, param)
}
