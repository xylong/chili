package middleware

import (
	"fmt"
	"github.com/xylong/chili"
)

// Authorization 认证
type Authorization struct {
}

func NewAuthorization() *Authorization {
	return &Authorization{}
}

func (a *Authorization) Before(ctx *chili.Context) error {
	fmt.Println("->auth")
	return nil
}

func (a *Authorization) After(data interface{}) (interface{}, error) {
	fmt.Println("auth->")
	return nil, nil
}
