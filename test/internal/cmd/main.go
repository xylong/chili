package main

import (
	"github.com/xylong/chili"
	v1 "github.com/xylong/chili/test/internal/api/v1"
)

type ReqTest struct {
	Name     string `json:"name" binding:"required"`     //用户名
	Password string `json:"password" binding:"required"` //新密码
}

func main() {
	chili.Ignite().
		Post("login", chili.GetHandleFunc(v1.Login)).
		Delete("logoff", chili.GetHandleFunc(v1.Logoff)).
		Launch()
}
