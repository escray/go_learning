package go_learning

import (
	"testing"
)

// func init() {
// 	fmt.Println("init1")
// }

// func init() {
// 	fmt.Println("init2")
// }

func TestPackage(t *testing.T) {
	t.Log(GetFibonacciSerie(5))
}
