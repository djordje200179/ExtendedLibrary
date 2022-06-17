package priorityqueue

type item[T any] struct {
	value    T
	priority int
}

type heapSlice[T any] struct {
	slice    []item[T]
	priority Priority
}

func (h *heapSlice[T]) Len() int {
	return len(h.slice)
}

func (h *heapSlice[T]) Less(i, j int) bool {
	if h.priority == SmallerFirst {
		return h.slice[i].priority < h.slice[j].priority
	} else {
		return h.slice[i].priority > h.slice[j].priority
	}
}

func (h *heapSlice[T]) Swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
}

func (h *heapSlice[T]) Push(value any) {
	h.slice = append(h.slice, value.(item[T]))
}

func (h *heapSlice[T]) Pop() any {
	sliceLen := h.Len()
	value := h.slice[sliceLen-1]
	h.slice = h.slice[0 : sliceLen-1]
	return value
}
