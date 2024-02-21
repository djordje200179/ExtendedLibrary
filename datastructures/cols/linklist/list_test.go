package linklist

import (
	"github.com/djordje200179/extendedlibrary/streams"
	"testing"
)

func TestListAppend(t *testing.T) {
	list := New[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	for i := range streams.From(list.Stream).Filter(func(x int) bool { return x%2 == 0 }) {
		if i != 2 {
			t.Errorf("Expected 2, got %d", i)
		}
	}
}
