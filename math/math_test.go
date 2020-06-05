package math

import (
	"testing"
)


func TestSum(t *testing.T)  {
	sum := Sum([]int {1,3,5})	
	if sum != 9 {
		t.Errorf("Wanted 9 but %d", sum)
	}
}



