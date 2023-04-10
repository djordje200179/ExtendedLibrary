package hashmap

import "unsafe"

type emptyInterface struct {
	typ unsafe.Pointer
	val unsafe.Pointer
}

type hmap struct{}

type maptype struct{}

func mapTypeAndValue(m interface{}) (*maptype, *hmap) {
	ei := (*emptyInterface)(unsafe.Pointer(&m))
	return (*maptype)(ei.typ), (*hmap)(ei.val)
}

//go:linkname internalMapGet runtime.mapaccess2
func internalMapGet(t *maptype, h *hmap, key unsafe.Pointer) (unsafe.Pointer, bool)
