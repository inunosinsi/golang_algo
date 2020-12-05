package tree

import (
	"testing"
)

// Measure the depth of a tree
func TestMeasureDepth(t *testing.T) {
	tests := []struct {
		values []int
		height int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{5, 4, 3, 2, 1}, 5},
		{[]int{9, 4, 15, 2, 6, 12, 17}, 3},
		{[]int{9, 4, 15, 2, 6, 12, 17, 1}, 4},
		{[]int{9, 4, 15, 2, 6, 12, 17, 1, 18}, 4},
		{[]int{9, 4, 15, 2, 6, 12, 17, 1, 18, 14}, 4},
	}

	for _, tt := range tests {
		tr := New(tt.values[0])
		for idx, i := range tt.values {
			if idx == 0 {
				continue
			}
			tr.Add(i)
		}
		if tt.height != MeasureDepth(tr) {
			t.Fatalf("failed")
		}
	}
}
