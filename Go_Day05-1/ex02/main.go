package main

import (
	"container/heap"
	"errors"
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
	return h[i].Size < h[j].Size
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
func getNCoolestPresents(h []Present, n int) ([]Present, error) {
	if len(h) < n || n <= 0 {
		return nil, errors.New("invalid value for n")
	}

	f := &Presents{}
	heap.Init(f)
	for _, k := range h {
		heap.Push(f, k)
	}

	var res []Present
	for f.Len() > 0 && len(res) < n {
		temp := heap.Pop(f).(Present)
		res = append(res, temp)
		next := (*f)[0]
		if temp.Value != next.Value {
			break
		}
	}
	return res, nil
}
