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
	mappedData := mappedData[KeyOut, ValueOut]{
		keyComparator: keyComparator,
	}

	for {
		var entry misc.Pair[KeyIn, ValueIn]
		entry, ok := <-dataSource

		if !ok {
			break
		}

		mapper(entry.First, entry.Second, mappedData.Append)
	}

	sort.Sort(&mappedData)

	var uniqueKeys []KeyOut
	var combinedValues []ValueOut

	lastIndex := -1
	for i := 1; i <= mappedData.Len(); i++ {
		lastKey := mappedData.keys[i-1]

		if i != mappedData.Len() {
			currentKey := mappedData.keys[i]

			if keyComparator(lastKey, currentKey) == comparison.Equal {
				continue
			}
		}

		firstIndex := lastIndex + 1
		lastIndex = i - 1

		validValues := mappedData.values[firstIndex : lastIndex+1]
		reducedValue := combiner(lastKey, validValues)

		uniqueKeys = append(uniqueKeys, lastKey)
		combinedValues = append(combinedValues, reducedValue)
	}

	appendData(uniqueKeys, combinedValues)

	finishSignal.Done()
}
