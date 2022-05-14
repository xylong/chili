package chili

import "fmt"

func BindJson[T any](resolve func(t *T, ctx *Context), reject func(t *T, err error, ctx *Context)) func(ctx *Context) {
	return func(ctx *Context) {
		var t T
		if err := ctx.ShouldBind(&t); err != nil {
			fmt.Println(err)
			reject(&t, err, ctx)
		} else {
			resolve(&t, ctx)
		}
	}
}
