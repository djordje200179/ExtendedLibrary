package linkedlist

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
)

func (list *List[T]) Sort(comparator functions.Comparator[T]) {
	mergeSort(list, comparator)
}

func mergeSort[T any](list *List[T], comparator functions.Comparator[T]) {
	if list == nil || list.head == nil || list.head.next == nil {
		return
	}

	firstList, secondList := splitList(list)

	mergeSort(firstList, comparator)
	mergeSort(secondList, comparator)

	sortedMerge(firstList, secondList, list, comparator)
}

func sortedMerge[T any](firstList, secondList, resultList *List[T], comparator functions.Comparator[T]) {
	if firstList == nil {
		*resultList = *secondList
	}

	if secondList == nil {
		*resultList = *firstList
	}

	resultList.Clear()

	firstListCurr := firstList.head
	secondListCurr := secondList.head

	for firstListCurr != nil && secondListCurr != nil {
		if comparator(firstListCurr.Value, secondListCurr.Value) == comparison.FirstSmaller {
			resultList.Append(firstListCurr.Value)
			firstListCurr = firstListCurr.next
		} else {
			resultList.Append(secondListCurr.Value)
			secondListCurr = secondListCurr.next
		}
	}

	for firstListCurr != nil {
		resultList.Append(firstListCurr.Value)
		firstListCurr = firstListCurr.next
	}

	for secondListCurr != nil {
		resultList.Append(secondListCurr.Value)
		secondListCurr = secondListCurr.next
	}
}

func splitList[T any](list *List[T]) (*List[T], *List[T]) {
	if list == nil || list.Size() <= 1 {
		return list, nil
	}

	middleIndex := list.Size() / 2
	middleNode := list.GetNode(middleIndex - 1)

	firstList := &List[T]{
		head: list.head,
		tail: middleNode,
		size: middleIndex,
	}

	secondList := &List[T]{
		head: middleNode.next,
		tail: list.tail,
		size: list.Size() - middleIndex,
	}

	middleNode.next.prev = nil
	middleNode.next = nil

	return firstList, secondList
}
