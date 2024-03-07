package array

import "testing"

func TestAppend(t *testing.T) {
	arr := FromSlice([]int{1, 2, 3})

	arr.Prepend(0)
	arr.Append(4)

	for i, val := range arr.Stream2 {
		t.Log(i, val)
	}
}
