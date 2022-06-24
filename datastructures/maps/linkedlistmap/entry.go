package linkedlistmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"
	"github.com/djordje200179/extendedlibrary/misc"
)

type entry[K comparable, V any] struct {
	node *linkedlist.Node[misc.Pair[K, V]]
}

func (e entry[K, V]) Key() K   { return e.node.Value.First }
func (e entry[K, V]) Value() V { return e.node.Value.Second }

func (e entry[K, V]) SetValue(value V) { e.node.Value.Second = value }
