package middleware

import (
	"fmt"
	"github.com/xylong/chili"
)

// Permission 访问权限
type Permission struct {
}

func NewPermission() *Permission {
	return &Permission{}
}

func (p *Permission) Before(ctx *chili.Context) error {
	fmt.Println("->p")
	return nil
}

func (p *Permission) After(data interface{}) (interface{}, error) {
	fmt.Println("p->")
	return nil, nil
}
