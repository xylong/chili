package chili

type (
	// Attrs 属性方法集合
	Attrs []Attr
	// Attr 属性方法
	Attr func(interface{})
	// ModelBuilder 创建模型方法
	ModelBuilder func(...Attr) interface{}
)

// Apply 设置属性
func (a Attrs) Apply(i interface{}) {
	for _, attr := range a {
		attr(i)
	}
}
