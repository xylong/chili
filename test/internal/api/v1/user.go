package v1

import (
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
func (c *UserController) ShowAccount(ctx *chili.Context) string {
	id := ctx.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		return id
	}

	return "hello"
}

func (c *UserController) Register(ctx *chili.Context) string {
	return "foo"
}
