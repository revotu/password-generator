package internal

import (
	"fmt"
	"testing"
)

func Test_RandNumber(t *testing.T) {
	for i := 0; i < 100; i++ {
		number := RandNumber(100)
		if number < 0 || number >= 100 {
			t.Fail()
		} else {
			fmt.Println(number)
		}
	}
}