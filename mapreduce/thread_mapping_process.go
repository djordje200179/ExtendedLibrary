package mapreduce

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"sort"
)

type threadMappingProcess[KeyIn any, ValueIn any, KeyOut comparable, ValueOut any] struct {
	keyComparator functions.Comparator[KeyOut]

	mapper   Mapper[KeyIn, ValueIn, KeyOut, ValueOut]
	combiner Reducer[KeyOut, ValueOut]

	mappedDataKeys   []KeyOut
	mappedDataValues []ValueOut
}

func (process *threadMappingProcess[KeyIn, ValueIn, KeyOut, ValueOut]) mapData(dataSource Source[KeyIn, ValueIn]) {
	mappedDataAppender := func(key KeyOut, value ValueOut) {
		process.mappedDataKeys = append(process.mappedDataKeys, key)
		process.mappedDataValues = append(process.mappedDataValues, value)
	}

	for {
		var entry misc.Pair[KeyIn, ValueIn]
		entry, ok := <-dataSource

		if !ok {
			break
		}

		process.mapper(entry.First, entry.Second, mappedDataAppender)
	}
}

func (process *threadMappingProcess[KeyIn, ValueIn, KeyOut, ValueOut]) sortData() {
	comparator := func(i, j int) bool {
		return process.keyComparator(process.mappedDataKeys[i], process.mappedDataKeys[j]) == comparison.FirstSmaller
	}

	sort.SliceStable(process.mappedDataKeys, comparator)
	sort.SliceStable(process.mappedDataValues, comparator)
}

func (process *threadMappingProcess[KeyIn, ValueIn, KeyOut, ValueOut]) combineData() {
	var uniqueKeys []KeyOut
	var combinedValues []ValueOut

	lastIndex := -1
	for i := 1; i <= len(process.mappedDataKeys); i++ {
		lastKey := process.mappedDataKeys[i-1]

		if i != len(process.mappedDataKeys) {
			currentKey := process.mappedDataKeys[i]

			if process.keyComparator(lastKey, currentKey) == comparison.Equal {
				continue
			}
		}

		firstIndex := lastIndex + 1
		lastIndex = i - 1

		validValues := process.mappedDataValues[firstIndex : lastIndex+1]
		reducedValue := process.combiner(lastKey, validValues)

		uniqueKeys = append(uniqueKeys, lastKey)
		combinedValues = append(combinedValues, reducedValue)
	}

	process.mappedDataKeys = uniqueKeys
	process.mappedDataValues = combinedValues
}
