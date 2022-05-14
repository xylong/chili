package chili

// Bind 绑定
type Bind[T any] struct {
	try   func(ctx *Context, t *T)
	catch func(ctx *Context, t *T, err error)
}

func NewBind[T any]() *Bind[T] {
	return &Bind[T]{}
}

func (b *Bind[T]) Try(f func(ctx *Context, t *T)) *Bind[T] {
	b.try = f
	return b
}

func (b *Bind[T]) Cache(f func(ctx *Context, t *T, err error)) *Bind[T] {
	b.catch = f
	return b
}

func (b *Bind[T]) Complete() func(ctx *Context) {
	return func(ctx *Context) {
		var t T
		if err := ctx.ShouldBind(&t); err != nil {
			b.catch(ctx, &t, err)
		} else {
			b.try(ctx, &t)
		}
	}
}
