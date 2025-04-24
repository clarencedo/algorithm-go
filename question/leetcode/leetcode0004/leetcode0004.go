package leetcode

import (
	"container/heap"
	"sort"
)

type maxHeap struct{ sort.IntSlice }

func (h maxHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}
func (h *maxHeap) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *maxHeap) Pop() any {
	n := len(h.IntSlice)
	v := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return v
}

type minHeap struct {
	sort.IntSlice
}

func (h minHeap) Less(i, j int) bool {
	return h.IntSlice[i] < h.IntSlice[j]
}
func (h *minHeap) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *minHeap) Pop() any {
	n := len(h.IntSlice)
	v := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return v
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	left := &maxHeap{}
	right := &minHeap{}
	heap.Init(left)
	heap.Init(right)

	for _, num := range append(nums1, nums2...) {
		if left.Len() == 0 || num <= left.IntSlice[0] {
			heap.Push(left, num)
		} else {
			heap.Push(right, num)
		}

		//balance heaps
		if left.Len() > right.Len()+1 {
			heap.Push(right, heap.Pop(left))
		}
		if right.Len() > left.Len() {
			heap.Push(left, heap.Pop(right))
		}
	}

	if left.Len() > right.Len() {
		return float64(left.IntSlice[0])
	}

	return float64(left.IntSlice[0]+right.IntSlice[0]) / 2.0
}