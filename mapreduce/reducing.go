package mapreduce

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"log"
	"sync"
)

func (process *Process[KeyIn, ValueIn, KeyOut, ValueOut]) reduceData() {
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
			if process.finalizer != nil {
				reducedValue = process.finalizer(lastKey, reducedValue)
			}

			process.mutex.Lock()
			_, err := fmt.Fprintf(process.dataWriter, "%v: %v\n", lastKey, reducedValue)
			if err != nil {
				log.Panic(err)
			}
			process.mutex.Unlock()

			barrier.Done()
		}()
	}

	barrier.Wait()
}
