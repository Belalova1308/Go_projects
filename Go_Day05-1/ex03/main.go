package main

import (
	"container/heap"
)

type Present struct {
	Value int
	Size  int
}

type Presents []Present

func (h Presents) Len() int { return len(h) }

func (h Presents) Less(i, j int) bool {
	if h[i].Value != h[j].Value {
		return h[i].Value > h[j].Value
	}
	return h[i].Size > h[j].Size
}
func (h Presents) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Presents) Push(x any) {
	*h = append(*h, x.(Present))
}
func (h *Presents) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func grabPresents(h []Present, capacity int) []Present {
	var res []Present
	f := &Presents{}
	heap.Init(f)
	for i := 0; i < len(h); i++ {
		heap.Push(f, h[i])
	}
	for capacity > 0 && f.Len() > 0 {
		val := heap.Pop(f).(Present)
		capacity -= val.Size
		res = append(res, val)
	}
	return res
}
