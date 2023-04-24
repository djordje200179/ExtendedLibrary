package mapreduce

import (
	"sync"
)

func reduceData[KeyOut, ValueOut any](
	reducer Reducer[KeyOut, ValueOut], finalizer Finalizer[KeyOut, ValueOut],
	write func(key KeyOut, value ValueOut), finishSignal *sync.WaitGroup,
	key KeyOut, values []ValueOut,
) {
	reducedValue := reducer(key, values)
	if finalizer != nil {
		reducedValue = finalizer(key, reducedValue)
	}

	write(key, reducedValue)

	finishSignal.Done()
}