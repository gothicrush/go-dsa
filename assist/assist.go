package assist

type Comparator func (i interface{}, j interface{}) int

func Compare(i interface{}, j interface{}, f func (i interface{}, j interface{}) int ) int {
	return f(i,j)
}