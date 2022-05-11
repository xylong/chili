package v1

import (
	"github.com/xylong/chili"
	"github.com/xylong/chili/test/internal/model"
	"github.com/xylong/chili/test/internal/model/user"
	"strconv"
)

type UserController struct {
}

// @Summary 详情
// @Description 用户详情
// @Tags 用户
// @Produce  json
// @Param id path string true "用户🆔"
// @Success 200 {object} user.User "success"
// @Router /users/{id} [get]
func (c *UserController) ShowAccount(ctx *chili.Context) chili.Json {
	id, _ := strconv.Atoi(ctx.Param("id"))
	return model.Model(model.User)(user.WithID(id)).(*user.User)
}

func (c *UserController) Register(ctx *chili.Context) string {
	return "foo"
}
