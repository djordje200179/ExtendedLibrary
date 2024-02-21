package linklist

import "testing"

func TestListAppend(t *testing.T) {
	list := New[int]()

	list.Append(1)
	list.Append(2)

	if list.Size() != 2 {
		t.Errorf("List size is not 2")
	}
}
