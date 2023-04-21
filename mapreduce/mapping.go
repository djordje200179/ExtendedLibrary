package mapreduce

import "sync"

func (process *Process[K, V]) appendData(key K, value V) {
	process.mutex.Lock()
	defer process.mutex.Unlock()

	process.mappedDataKeys = append(process.mappedDataKeys, key)
	process.mappedDataValues = append(process.mappedDataValues, value)
}

func (process *Process[K, V]) mapData() {
	var barrier sync.WaitGroup
	barrier.Add(len(process.dataSources))

	for _, dataSource := range process.dataSources {
		dataSource := dataSource
		go func() {
			dataSource.Map(process.appendData)
			barrier.Done()
		}()
	}

	barrier.Wait()
}
