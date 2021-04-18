package go_paradigms

import (
	"bytes"
	"fmt"
	"testing"
	"unsafe"
)

type T struct {
	name  string // name of the object
	value int    // its value
}

type slice struct {
	// 指向存放数据的数组指针
	array unsafe.Pointer
	// 长度有多大
	len int
	// 容量有多大
	cap int
}

func TestSliceShare(t *testing.T) {
	foo := make([]int, 5)
	sliceInfo(foo)
	foo[3] = 42
	foo[4] = 100
	bar := foo[1:4]
	bar[1] = 99
	sliceInfo(foo)
	sliceInfo(bar)
}

func TestSliceAppend(t *testing.T) {
	a := make([]int, 32)
	b := a[0:15]
	sliceInfo(a)
	sliceInfo(b)
	a = append(a, 1)
	a[2] = 42
	sliceInfo(a)
	sliceInfo(b)
}

func TestSliceAppendAgain(t *testing.T) {
	path := []byte("AAAA/BBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	// dir1 := path[:sepIndex]
	dir1 := path[:sepIndex:sepIndex]
	dir2 := path[sepIndex+1:]
	arrayInfo(dir1)
	arrayInfo(dir2)
	dir1 = append(dir1, "suffix"...)
	arrayInfo(dir1)
	arrayInfo(dir2)
}

func sliceInfo(slice []int) {
	fmt.Println(slice, len(slice), cap(slice), unsafe.Pointer(&slice))
}

func arrayInfo(arr []byte) {
	fmt.Println(string(arr), unsafe.Pointer(&arr))
}
