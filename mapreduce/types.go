package mapreduce

import "github.com/djordje200179/extendedlibrary/misc"

type Source[KeyIn any, ValueIn any] <-chan misc.Pair[KeyIn, ValueIn]
type Emitter[KeyOut comparable, ValueOut any] func(key KeyOut, value ValueOut)
type Mapper[KeyIn any, ValueIn any, KeyOut comparable, ValueOut any] func(key KeyIn, value ValueIn, emit Emitter[KeyOut, ValueOut])
type Reducer[KeyOut comparable, ValueOut any] func(key KeyOut, values []ValueOut) ValueOut
type Finalizer[KeyOut comparable, ValueOut any] func(key KeyOut, value ValueOut) ValueOut
