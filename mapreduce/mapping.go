package mapreduce

import (
	"runtime"
	"sync"
)

func (process *Process[K, V]) mapData() {
	var barrier sync.WaitGroup

	threadsCount := runtime.NumCPU()
	sourcesPerThread := len(process.dataSources) / threadsCount

	threadProcesses := make([]threadProcess[K, V], threadsCount)
	for i := 0; i < threadsCount; i++ {
		firstSourceIndex := i * sourcesPerThread
		lastSourceIndex := (i + 1) * sourcesPerThread
		sources := process.dataSources[firstSourceIndex:lastSourceIndex]

		threadProcesses[i] = threadProcess[K, V]{
			dataSources:   sources,
			keyComparator: process.keyComparator,
			combiner:      process.reducer,
		}

		currProcess := &threadProcesses[i]

		barrier.Add(1)

		go func() {
			currProcess.mapData()
			currProcess.sortData()
			currProcess.combineData()

			process.mutex.Lock()
			process.mappedDataKeys = append(process.mappedDataKeys, currProcess.mappedDataKeys...)
			process.mappedDataValues = append(process.mappedDataValues, currProcess.mappedDataValues...)
			process.mutex.Unlock()

			barrier.Done()
		}()
	}

	barrier.Wait()
}
