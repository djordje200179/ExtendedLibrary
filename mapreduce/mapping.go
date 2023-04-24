package mapreduce

import (
	"runtime"
	"sync"
)

func (process *Process[KeyIn, ValueIn, KeyOut, ValueOut]) mapData() {
	threadsCount := runtime.NumCPU()

	var barrier sync.WaitGroup
	barrier.Add(threadsCount)

	for i := 0; i < threadsCount; i++ {
		go mapData(
			process.keyComparator,
			process.mapper, process.reducer,
			process.dataSource,
			process.appendData, &barrier,
		)
	}

	barrier.Wait()
}

func (process *Process[KeyIn, ValueIn, KeyOut, ValueOut]) appendData(keys []KeyOut, values []ValueOut) {
	process.mutex.Lock()
	process.mappedData.keys = append(process.mappedData.keys, keys...)
	process.mappedData.values = append(process.mappedData.values, values...)
	process.mutex.Unlock()
}
