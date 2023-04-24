package mapreduce

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"golang.org/x/exp/constraints"
	"io"
	"sort"
	"sync"
)

type Process[KeyIn, ValueIn, KeyOut, ValueOut any] struct {
	keyComparator functions.Comparator[KeyOut]

	mapper    Mapper[KeyIn, ValueIn, KeyOut, ValueOut]
	reducer   Reducer[KeyOut, ValueOut]
	finalizer Finalizer[KeyOut, ValueOut]

	dataSource Source[KeyIn, ValueIn]

	mutex sync.Mutex

	mappedData mappedData[KeyOut, ValueOut]

	dataWriter   io.Writer
	finishSignal sync.WaitGroup
}

func NewProcess[KeyIn, ValueIn, KeyOut, ValueOut any](
	keyComparator functions.Comparator[KeyOut],
	mapper Mapper[KeyIn, ValueIn, KeyOut, ValueOut], reducer Reducer[KeyOut, ValueOut], finalizer Finalizer[KeyOut, ValueOut],
	output io.Writer,
	dataSource Source[KeyIn, ValueIn],
) *Process[KeyIn, ValueIn, KeyOut, ValueOut] {
	//output = bufio.NewWriter(output)

	process := &Process[KeyIn, ValueIn, KeyOut, ValueOut]{
		keyComparator: keyComparator,

		mapper:    mapper,
		reducer:   reducer,
		finalizer: finalizer,

		dataSource: dataSource,

		mappedData: mappedData[KeyOut, ValueOut]{
			keyComparator: keyComparator,
		},

		dataWriter: output,
	}

	return process
}

func NewProcessWithOrderedKeys[KeyIn, ValueIn any, KeyOut constraints.Ordered, ValueOut any](
	mapper Mapper[KeyIn, ValueIn, KeyOut, ValueOut], reducer Reducer[KeyOut, ValueOut], finalizer Finalizer[KeyOut, ValueOut],
	output io.Writer,
	dataSource Source[KeyIn, ValueIn],
) *Process[KeyIn, ValueIn, KeyOut, ValueOut] {
	return NewProcess(comparison.Ascending[KeyOut], mapper, reducer, finalizer, output, dataSource)
}

func (process *Process[KeyIn, ValueIn, KeyOut, ValueOut]) Run() {
	process.mapData()
	sort.Sort(&process.mappedData)
	process.reduceData()

	process.finishSignal.Done()
}

func (process *Process[KeyIn, ValueIn, KeyOut, ValueOut]) WaitToFinish() {
	process.finishSignal.Add(1)
	process.finishSignal.Wait()
}
