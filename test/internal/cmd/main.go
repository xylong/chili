package main

import (
	"github.com/xylong/chili"
	v1 "github.com/xylong/chili/test/internal/api/v1"
)

func main() {
	chili.Ignite().
		Post("register", v1.Register).
		Post("login", v1.Login).
		Delete("logoff", v1.Logoff).
		Launch()
}
