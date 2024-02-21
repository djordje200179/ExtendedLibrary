package streams

func (s Stream2[K, V]) Count() int {
	count := 0
	for _, _ = range s {
		count++
	}

	return count
}
