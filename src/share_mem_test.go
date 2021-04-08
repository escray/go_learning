package go_learning

import (
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
		time.Sleep(1 * time.Second)
		t.Logf("counter = %d", counter)
	}
}
