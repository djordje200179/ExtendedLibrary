package streams

import "testing"

func TestEnumerate(t *testing.T) {
	for i, val := range Enumerate(Range(5, 9)) {
		t.Log(i, val)
	}
}
