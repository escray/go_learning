package string_test

import (
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	var s string
	// 初始化默认为零值 ""
	t.Log(s)
	s = "hello"
	t.Log(len(s)) // 5
	// string 是不可变的 byte slice
	// s[1] = '3'
	// 可以存储任何二进制数据
	s = "\xE4\xB8\xA5"
	t.Log(s) // 严
	s = "\xE4\xBA\xB5"
	t.Log(s) // 亵
	s = "中"
	// 是 byte 数
	t.Log(len(s)) // 3

	c := []rune(s)
	t.Log("rune size: ", unsafe.Sizeof(c[0])) // rune size:  4
	t.Logf("中 unicode %x", c[0]) // 中 unicode 4e2d
	t.Logf("中 UTF8 %x", s) // 中 UTF8 e4b8ad
}

func TestStringToRune(t *testing.T) {
	s := "我究竟想要找一个什么样的工作？"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}
}
