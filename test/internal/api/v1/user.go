package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
func (c *UserController) ShowAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	fmt.Println(aid)
	if err != nil {
		ctx.JSONP(400, gin.H{"code": 200, "data": nil, "msg": "ok"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": "foo", "msg": "ok"})
}

func (c *UserController) Register(ctx *gin.Context) {
	ctx.JSONP(200, gin.H{"code": 200, "data": nil, "msg": "ok"})
}
