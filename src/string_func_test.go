package string_test

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))

	slice := []rune("string")
	t.Log(slice) // [115 116 114 105 110 103]
	za := []rune("早安")
	t.Log(za) // [26089 23433]

	ss := "\xe4\xb8\x2d"
	r := []rune(ss)
	t.Log(len(ss))
	t.Log(r)	// [65533 65533 45]
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}

func TestStringHeader(t *testing.T) {
	str1 := "hello world"
	str2 := "hello world"
	header1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	header2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))
	t.Log(header1.Data == header2.Data) // true
	t.Log(header1 == header2) // false
	t.Log(unsafe.Pointer(&str1), unsafe.Pointer(&str2)) // 0xc000096530 0xc000096540
}
