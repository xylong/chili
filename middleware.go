package chili

type middlewares []middleware

// middleware 中间件
type middleware interface {
	// Before 执行路由方法前调用
	Before(ctx *Context) error

	// After 执行路由方法后调用
	After(interface{}) (interface{}, error)
}

func (m middlewares) before(ctx *Context) {
	for _, f := range m {
		if err := f.Before(ctx); err != nil {
			panic(err)
		}
	}
}
