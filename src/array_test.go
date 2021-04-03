package go_learning

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Log(arr[1], arr[2]) // default 0
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	for idx, e := range arr {
		t.Log(idx, e)
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	arr_sec := arr[:3]
	t.Log(arr_sec)
	t.Log(arr[:2])
	t.Log(arr[:])
}

func TestSlice(t *testing.T) {
	var slice []int
	t.Log(len(slice), cap(slice))	// 0 0
	slice = append(slice, 1)
	t.Log(len(slice), cap(slice)) // 1 1

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))			// 4 4
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))			// 3 5
	t.Log(s2[0], s2[1], s2[2])  // 0 0 0
	// panic: runtime error: index out of range [3] with length 3
	// t.Log(s2[0], s2[1], s2[2], s2[3])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3]) // 0 0 0 1
	t.Log(len(s2), cap(s2))			// 4 5
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		// append(s, i)
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}
// 1 1
// 2 2
// 3 4
// 4 4
// 5 8
// 6 8
// 7 8
// 8 8
// 9 16
// 10 16

func TestSliceGrowth(t *testing.T) {
	s3 := make([]int, 1)
	var n int
	for n < 1500 {
		t.Log(n, len(s3), cap(s3))
		s3 = append(s3, 1)
		n++
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2)) // [Apr May Jun] 3 9
	Summer := year[5:8]
	t.Log(Summer, len(Summer), cap(Summer)) // [Jun Jul Aug] 3 7
	Summer[0] = "Unknow"
	t.Log(Q2) // [Apr May Unknow]
	t.Log(year) // [Jan Feb Mar Apr May Unknow Jul Aug Sep Oct Nov Dec]
}

func TestSliceCamparing(t *testing.T) {
	// a := []int{1, 2, 3, 4}
	// b := []int{1, 2, 3, 4}

	// invalid operation: a == b (slice can only be compared to nil)
	// if a == b {
	// 	t.Log("equal")
	// }
}

func TestSliceCompare(t *testing.T) {
	var a []int
	var b = make([]int, 0, 0)
	c := []int{}
	t.Log(a, len(a), cap(a))
	t.Log(b, len(b), cap(b))
	t.Log(c, len(c), cap(c))
	t.Log(a == nil, b == nil, c == nil)
}

func TestSliceGrows(t *testing.T) {
	a := []int64{}
	t.Log(a, len(a), cap(a)) // [] 0 0
	a = append(a, 1, 2, 3, 4, 5, 6, 7)
	t.Log(a, len(a), cap(a)) // [1 2 3 4 5 6 7] 7 8

	b := make([]int, 0, 5)
	t.Log(b, len(b), cap(b)) // [] 0 5
	b = append(b, 1, 2, 3, 4, 5)
	t.Log(b, len(b), cap(b)) // [1 2 3 4 5] 5 5
	b = append(b, 6)
	t.Log(b, len(b), cap(b)) // [1 2 3 4 5 6] 6 10
}
