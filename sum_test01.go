package calc

import "testing"

func TestSum(t *testing.T) {
	r := Sum(1,2)
	if r != 3 {
		t.Error("It is not 3. r=", r)
	}
}
