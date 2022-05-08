package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xylong/chili"
	"strconv"
)

type UserController struct {
}

// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {string}  gin.H{"code":200,"data":null,"msg":"ok"}
// @Router       /users/{id} [get]
func (c *UserController) ShowAccount(ctx *gin.Context) *chili.Response {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		return chili.R().Json(gin.H{"code": 200, "data": nil, "msg": "ok"})
	}

	return chili.R().Json(gin.H{"code": 200, "data": aid, "msg": "ok"})
}

func (c *UserController) Register(ctx *gin.Context) {
	ctx.JSONP(200, gin.H{"code": 200, "data": nil, "msg": "ok"})
}
