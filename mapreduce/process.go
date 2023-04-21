package mapreduce

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"golang.org/x/exp/constraints"
	"io"
	"log"
	"sort"
	"sync"
)

type Process[K comparable, V any] struct {
	mutex sync.Mutex

	dataSources []Mapper[K, V]

	mappedDataKeys   []K
	mappedDataValues []V
	keyComparator    functions.Comparator[K]

	reducer     Reducer[K, V]
	reducedData map[K]V

	finalizer    Finalizer[K, V]
	dataWriter   io.Writer
	finishSignal sync.WaitGroup
}

func NewProcess[K comparable, V any](
	keyComparator functions.Comparator[K],
	reducer Reducer[K, V], finalizer Finalizer[K, V],
	output io.Writer,
	dataSources ...Mapper[K, V],
) *Process[K, V] {
	process := &Process[K, V]{
		dataSources: dataSources,

		mappedDataKeys:   make([]K, 0),
		mappedDataValues: make([]V, 0),
		keyComparator:    keyComparator,

		reducer:     reducer,
		reducedData: make(map[K]V),

		finalizer: finalizer,

		dataWriter: output,
	}

	return process
}

func NewDefaultProcess[K constraints.Ordered, V any](
	reducer Reducer[K, V],
	output io.Writer,
	dataSources ...Mapper[K, V],
) *Process[K, V] {
	return NewProcess[K, V](comparison.Ascending[K], reducer, nil, output, dataSources...)
}

func (process *Process[K, V]) Run() {
	process.mapData()
	process.sortData()
	process.reduceData()
	if process.finalizer != nil {
		process.finalizeData()
	}
	process.outputData()

	process.finishSignal.Done()
}

func (process *Process[K, V]) WaitToFinish() {
	process.finishSignal.Add(1)
	process.finishSignal.Wait()
}

func (process *Process[K, V]) sortData() {
	keysComparator := func(i, j int) bool {
		return process.keyComparator(process.mappedDataKeys[i], process.mappedDataKeys[j]) == comparison.FirstSmaller
	}

	sort.SliceStable(process.mappedDataKeys, keysComparator)
	sort.SliceStable(process.mappedDataValues, keysComparator)
}

func (process *Process[K, V]) finalizeData() {
	for key, value := range process.reducedData {
		finalizedValue := process.finalizer(key, value)
		process.reducedData[key] = finalizedValue
	}
}

func (process *Process[K, V]) outputData() {
	for key, value := range process.reducedData {
		_, err := fmt.Fprintf(process.dataWriter, "%v: %v\n", key, value)
		if err != nil {
			log.Panic(err)
		}
	}
}
