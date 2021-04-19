package effective_go

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	x := []int{1,2,3}
	x = append(x, 4, 5, 6)
	fmt.Println(x)

	x = []int{1,2,3}
	y := []int{4,5,6}
	x = append(x, y...)
	fmt.Println(x)
}
