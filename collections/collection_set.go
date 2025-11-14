package Collections

// Set 泛型Set实现
type Set[T comparable] map[T]struct{}

// NewSet 创建新Set
func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

// Add 添加元素
func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

// Remove 删除元素
func (s Set[T]) Remove(value T) {
	delete(s, value)
}

// Contains 检查元素是否存在
func (s Set[T]) Contains(value T) bool {
	_, exists := s[value]
	return exists
}

// Size 获取Set大小
func (s Set[T]) Size() int {
	return len(s)
}

// ToSlice 转换为切片
func (s Set[T]) ToSlice() []T {
	keys := make([]T, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
