package assist

// 自定义比较器
type Comparator func (i interface{}, j interface{}) int


// 比较函数，第1,2个参数接收需要比较的值，第3个参数传入自定义的比较器
func Compare(i interface{}, j interface{}, f func (i interface{}, j interface{}) int ) int {
	return f(i,j)
}

// 自定义访问器
type Visitor func (v interface{}) 