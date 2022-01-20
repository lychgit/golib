package golib

import (
"fmt"
"testing"
)

func TestInArray(t *testing.T)  {
	var a = []int{
		1,
		2,
	}
	res := InArray(a, 1)
	fmt.Printf("res: %t \n", res)
}
