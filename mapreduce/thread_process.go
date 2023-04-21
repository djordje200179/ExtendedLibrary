package mapreduce

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"sort"
)

type threadProcess[K comparable, V any] struct {
	dataSources   []Mapper[K, V]
	keyComparator functions.Comparator[K]

	mappedDataKeys   []K
	mappedDataValues []V

	combiner Reducer[K, V]
}

func (process *threadProcess[K, V]) mapData() {
	mappedDataAppender := func(key K, value V) {
		process.mappedDataKeys = append(process.mappedDataKeys, key)
		process.mappedDataValues = append(process.mappedDataValues, value)
	}

	for _, dataSource := range process.dataSources {
		dataSource.Map(mappedDataAppender)
	}
}

func (process *threadProcess[K, V]) sortData() {
	comparator := func(i, j int) bool {
		return process.keyComparator(process.mappedDataKeys[i], process.mappedDataKeys[j]) == comparison.FirstSmaller
	}

	sort.SliceStable(process.mappedDataKeys, comparator)
	sort.SliceStable(process.mappedDataValues, comparator)
}

func (process *threadProcess[K, V]) combineData() {
	var uniqueKeys []K
	var uniqueValues []V

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
		uniqueValues = append(uniqueValues, reducedValue)
	}

	process.mappedDataKeys = uniqueKeys
	process.mappedDataValues = uniqueValues
}
