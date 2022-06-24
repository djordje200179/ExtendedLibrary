package linkedlistmap

import (
	"github.com/djordje200179/extendedlibrary/misc"
)

type entry[K comparable, V any] struct {
	pair *misc.Pair[K, V]
}

func (e entry[K, V]) Key() K { return e.pair.First }

func (e entry[K, V]) ValueRef() *V     { return &e.pair.Second }
func (e entry[K, V]) Value() V         { return *e.ValueRef() }
func (e entry[K, V]) SetValue(value V) { *e.ValueRef() = value }
