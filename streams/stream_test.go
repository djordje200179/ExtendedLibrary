package streams

import "testing"

func TestEnumerate(t *testing.T) {
	for i, val := range Range(5, 9).Enumerate() {
		t.Log(i, val)
	}
}
