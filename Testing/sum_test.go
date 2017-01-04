package addr

import "testing"

type sumTest struct {
	in1, in2, out int
}

var sumTests = []sumTest{
	sumTest{1, 1, 2},
	sumTest{2, 4, 6},
}

func TestSum(t *testing.T) {
	for _, dt := range sumTests {
		v := Sum(dt.in1, dt.in2)
		if v != dt.out {
			t.Error("Sum(%d, %d) = %d, want %d.", dt.in1, dt.in2, v, dt.out)
		}
	}
}
