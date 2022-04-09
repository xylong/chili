package middleware

import (
	"fmt"
	"github.com/xylong/chili"
)

// Csrf 跨域
type Csrf struct {
}

func NewCsrf() *Csrf {
	return &Csrf{}
}

func (c *Csrf) Before(ctx *chili.Context) error {
	fmt.Println("->csrf")
	return nil
}

func (c *Csrf) After(data interface{}) (interface{}, error) {
	fmt.Println("csrf->")
	return nil, nil
}
