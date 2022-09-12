package main

type Heap struct {
	Buckets  []int
	Capacity int
	Count    int
}

func NewHeap(capactiy int) *Heap {
	return &Heap{
		Buckets:  make([]int, capactiy+1),
		Capacity: capactiy + 1,
		Count:    0,
	}
}

func (h *Heap) Insert(data int) {
	if h.Count > h.Capacity {
		return
	}
	h.Count++
	h.Buckets[h.Count] = data
	i := h.Count
	for {
		if i/2 > 0 && h.Buckets[i] > h.Buckets[i/2] { //自下而上的堆化
			h.Buckets[i/2], h.Buckets[i] = h.Buckets[i], h.Buckets[i/2]
		}
		i /= 2
	}
}

func (h *Heap) Remove() int {
	if h.Count == 0 {
		return -1
	}
	value := h.Buckets[1]
	h.Buckets[1] = h.Buckets[h.Count]
	h.Count--
	heapfy(h.Buckets, h.Count, 1)
	return value
}

func heapfy(heap []int, n, i int) { //自上而下的堆化
	for {
		mostMax := i
		if 2*i <= n && heap[i] < heap[2*i] {
			mostMax = 2 * i
		}
		if 2*i+1 <= n && heap[mostMax] < heap[2*i+1] {
			mostMax = 2*i + 1
		}
		if mostMax == i {
			break
		}
		heap[i], heap[mostMax] = heap[mostMax], heap[i]
		i = mostMax
	}
}

func buildHeap(heap []int, n int) {
	for i := n / 2; i >= 1; i-- {
		heapfy(heap, n, i)
	}
}

func heapSort(heap []int, n int) {
	buildHeap(heap, n)
	k := n
	for k > 1 {
		heap[1], heap[k] = heap[k], heap[1]
		k--
		heapfy(heap, k, 1)
	}
}
