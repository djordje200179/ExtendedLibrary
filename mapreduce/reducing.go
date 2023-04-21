package mapreduce

import (
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"sync"
)

func (process *Process[K, V]) reduceData() {
	var barrier sync.WaitGroup

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
		barrier.Add(1)

		go func() {
			reducedValue := process.reducer(lastKey, validValues)

			process.mutex.Lock()
			process.reducedData[lastKey] = reducedValue
			process.mutex.Unlock()

			barrier.Done()
		}()
	}

	barrier.Wait()
}
