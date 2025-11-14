package Collections

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	// 创建整数Set
	intSet := NewSet[int]()
	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(1) // 重复添加不会生效

	fmt.Println("Integer Set size:", intSet.Size())

	// 创建字符串Set
	stringSet := NewSet[string]()
	stringSet.Add("hello")
	stringSet.Add("world")
	stringSet.Add("hello") // 重复添加不会生效

	fmt.Println("String Set size:", stringSet.Size())
}
