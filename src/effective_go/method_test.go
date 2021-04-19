package effective_go

import (
	"fmt"
	"testing"
)

type ByteSlice []byte

// func (slice ByteSlice) Append(data []byte) []byte{
// 	l := len(slice)
// 	if l + len(data) > cap(slice) {
// 		newSlice := make([]byte, (l+len(data))*2)
// 		copy(newSlice, slice)
// 		slice = newSlice
// 	}
// 	slice = slice[0:l+len(data)]
// 	copy(slice[l:], data)
// 	return slice
// }

func (p *ByteSlice) Append(data []byte) {
	slice := *p
	l := len(slice)
	if l + len(data) > cap(slice) {
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l+len(data)]
	copy(slice[l:], data)
	*p = slice
}

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	l := len(slice)
	if l + len(data) > cap(slice) {
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l+len(data)]
	copy(slice[l:], data)
	*p = slice
	return len(data), nil
}

func TestByteSlicePointer(t *testing.T) {
	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
	fmt.Println(string(b))
}
