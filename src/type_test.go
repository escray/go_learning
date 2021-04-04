package go_learning

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T){
	var a int32 = 1
	var b int64 = int64(a)
	t.Log(a, b)

	var c MyInt = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr := aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*"+s+"*")
	t.Log(len(s))
}

func TestFromComments(t *testing.T) {
	cpu := runtime.GOARCH
	int_size := strconv.IntSize
	t.Log(cpu, int_size)
}

func TestConvert(t *testing.T) {
	i := int64(10)
	j := int8(i)
	fmt.Println(j)

	a := int64(math.MaxInt64)
	b := int32(a)
	t.Log(a, b)
}
