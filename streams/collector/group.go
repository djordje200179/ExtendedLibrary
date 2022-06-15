package collector

//func Group[T any, P comparable](stream Stream[T], mapper functions.Mapper[T, P]) Stream[misc.Pair[P, []T]] {
//	ret := create[misc.Pair[P, []T]]()
//
//	go func() {
//		if !ret.waitRequest() {
//			ret.close()
//			return
//		}
//
//		m := make(map[P][]T)
//		stream.ForEach(func(data T) {
//			key := mapper(data)
//
//			_, ok := m[key]
//			if !ok {
//				m[key] = []T{}
//			}
//
//			m[key] = append(m[key], data)
//		})
//
//		entries := make([]misc.Pair[P, []T], 0, len(m))
//		for key, value := range m {
//			entries = append(entries, misc.Pair[P, []T]{key, value})
//		}
//
//		ret.data <- entries[0]
//		for i := 1; i < len(entries) && ret.waitRequest(); i++ {
//			ret.data <- entries[i]
//		}
//
//		ret.close()
//	}()
//
//	return ret
//}
