package v1

import (
	"github.com/xylong/chili"
	"github.com/xylong/chili/test/internal/model"
	"github.com/xylong/chili/test/internal/model/user"
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
func (c *UserController) ShowAccount(ctx *chili.Context) *user.User {
	id, _ := strconv.Atoi(ctx.Param("id"))
	return model.Model(model.User)(user.WithID(id)).(*user.User)
}

func (c *UserController) Register(ctx *chili.Context) string {
	return "foo"
}
