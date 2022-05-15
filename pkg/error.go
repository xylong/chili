package pkg

// ReturnValue 返回值
type ReturnValue[T any] struct {
	Data T
	Err  error
}

// Unwrap 抛出错误
func (v *ReturnValue[T]) Unwrap() T {
	if v.Err != nil {
		panic(v.Err)
	}
	return v.Data
}

// Value 返回值
func Value[T any](value T, err error) *ReturnValue[T] {
	return &ReturnValue[T]{Data: value, Err: err}
}
