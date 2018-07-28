package module3

// Le 乐
type Le struct {
	name string
}

// NewLe 新建
func NewLe(name string) *Le {
	return &Le{
		name: name,
	}
}

// String 字符串
func (l Le) String() string {
	return l.name
}
