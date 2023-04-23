package mapreduce

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"sort"
	"sync"
)

func mapData[KeyIn, ValueIn, KeyOut, ValueOut any](
	keyComparator functions.Comparator[KeyOut],
	mapper Mapper[KeyIn, ValueIn, KeyOut, ValueOut], combiner Reducer[KeyOut, ValueOut],
	dataSource Source[KeyIn, ValueIn],
	appendData func(keys []KeyOut, values []ValueOut), finishSignal *sync.WaitGroup,
) {
	var mappedDataKeys []KeyOut
	var mappedDataValues []ValueOut

	mappedDataAppender := func(key KeyOut, value ValueOut) {
		mappedDataKeys = append(mappedDataKeys, key)
		mappedDataValues = append(mappedDataValues, value)
	}

	for {
		var entry misc.Pair[KeyIn, ValueIn]
		entry, ok := <-dataSource

		if !ok {
			break
		}

		mapper(entry.First, entry.Second, mappedDataAppender)
	}

	comparator := func(i, j int) bool {
		return keyComparator(mappedDataKeys[i], mappedDataKeys[j]) == comparison.FirstSmaller
	}

	sort.Slice(mappedDataValues, comparator)
	sort.Slice(mappedDataKeys, comparator)

	var uniqueKeys []KeyOut
	var combinedValues []ValueOut

	lastIndex := -1
	for i := 1; i <= len(mappedDataKeys); i++ {
		lastKey := mappedDataKeys[i-1]

		if i != len(mappedDataKeys) {
			currentKey := mappedDataKeys[i]

			if keyComparator(lastKey, currentKey) == comparison.Equal {
				continue
			}
		}

		firstIndex := lastIndex + 1
		lastIndex = i - 1

		validValues := mappedDataValues[firstIndex : lastIndex+1]
		reducedValue := combiner(lastKey, validValues)

		uniqueKeys = append(uniqueKeys, lastKey)
		combinedValues = append(combinedValues, reducedValue)
	}

	appendData(uniqueKeys, combinedValues)

	finishSignal.Done()
}
