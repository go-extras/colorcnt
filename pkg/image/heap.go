package image

import "container/heap"

type kv struct {
	Key   RGBColor
	Value int
}

type kvheap []kv

func (h kvheap) Len() int           { return len(h) }
func (h kvheap) Less(i, j int) bool { return h[i].Value > h[j].Value }
func (h kvheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *kvheap) Push(x interface{}) {
	*h = append(*h, x.(kv))
}

func (h *kvheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getHeap(m RGBColorMap) *kvheap {
	h := &kvheap{}
	heap.Init(h)
	for k, v := range m {
		heap.Push(h, kv{k, v})
	}
	return h
}
