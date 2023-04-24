package mapreduce

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
)

type mappedData[KeyOut, ValueOut any] struct {
	keyComparator functions.Comparator[KeyOut]

	keys   []KeyOut
	values []ValueOut
}

func (data *mappedData[KeyOut, ValueOut]) Append(key KeyOut, value ValueOut) {
	data.keys = append(data.keys, key)
	data.values = append(data.values, value)
}
func (data *mappedData[KeyOut, ValueOut]) Len() int {
	return len(data.keys)
}

func (data *mappedData[KeyOut, ValueOut]) Less(i, j int) bool {
	return data.keyComparator(data.keys[i], data.keys[j]) == comparison.FirstSmaller
}

func (data *mappedData[KeyOut, ValueOut]) Swap(i, j int) {
	data.keys[i], data.keys[j] = data.keys[j], data.keys[i]
	data.values[i], data.values[j] = data.values[j], data.values[i]
}
