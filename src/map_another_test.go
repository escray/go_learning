package map_ext

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	m := map[int]func(op int)int {}
	// Lambda 定义，形参可以省略
	// m := map[int]func(int)int {}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op*op }
	m[3] = func(op int) int { return op*op*op }
	t.Log(m[1](2), m[2](2), m[3](2)) // 2 4 8
	// 通过 key 访问其后的函数
}

func TestMapForset(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	setExisting(mySet, n)
	t.Log(len(mySet)) // 1
	mySet[3] = true
	t.Log(len(mySet)) // 2
	delete(mySet, 1)
	setExisting(mySet, 3)
	t.Log(len(mySet)) // 1
}

func setExisting(set map[int]bool, n int) {
	if set[n] {
		fmt.Printf("%d is existing", n)
	} else {
		fmt.Printf("%d is not existing", n)
	}
}

