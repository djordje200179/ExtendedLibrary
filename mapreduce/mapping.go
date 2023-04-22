package mapreduce

import (
	"runtime"
	"sync"
)

func (process *Process[KeyIn, ValueIn, KeyOut, ValueOut]) mapData() {
	threadsCount := runtime.NumCPU()

	var barrier sync.WaitGroup
	barrier.Add(threadsCount)

	threadProcesses := make([]threadMappingProcess[KeyIn, ValueIn, KeyOut, ValueOut], threadsCount)
	for i := range threadProcesses {
		threadProcesses[i] = threadMappingProcess[KeyIn, ValueIn, KeyOut, ValueOut]{
			keyComparator: process.keyComparator,
			mapper:        process.mapper,
			combiner:      process.reducer,
		}

		currProcess := &threadProcesses[i]

		go func() {
			currProcess.mapData(process.dataSource)
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
