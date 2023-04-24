package mapreduce

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
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

	uniqueKeys, combinedValues := mappedData.Reduce(combiner)

	appendData(uniqueKeys, combinedValues)

	finishSignal.Done()
}
