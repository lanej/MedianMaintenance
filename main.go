package main

import (
	"container/heap"
	"io/ioutil"
	"strconv"
	"strings"
)

// An MinHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push me
func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// Pop me
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// An MaxHeap is a max-heap of ints.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push me
func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// Pop me
func (h *MaxHeap) Pop() interface{} {
	heapSize := len(*h)
	lastNode := (*h)[heapSize-1]
	*h = (*h)[0 : heapSize-1]
	return lastNode
}

// Median hold the median
type Median struct {
	minHeap MinHeap
	maxHeap MaxHeap
}

// NewMedian intializes a new median
func NewMedian(numbers []int) *Median {
	median := &Median{
		minHeap: MinHeap{},
		maxHeap: MaxHeap{},
	}

	heap.Init(&median.minHeap)
	heap.Init(&median.maxHeap)

	for i := 0; i < len(numbers); i++ {
		median.Add(numbers[i])
	}

	return median
}

// Add a number to the running median
func (m *Median) Add(num int) {
	minHeapLen := m.minHeap.Len()
	maxHeapLen := m.maxHeap.Len()

	if minHeapLen == 0 && maxHeapLen == 0 {
		heap.Push(&m.maxHeap, num)
	} else if float64(num) > float64(m.maxHeap[0]) {
		heap.Push(&m.minHeap, num)
	} else {
		heap.Push(&m.maxHeap, num)
	}

	// rebalance
	if len(m.maxHeap)-len(m.minHeap) > 1 {
		heap.Push(&m.minHeap, heap.Pop(&m.maxHeap))
	} else if len(m.minHeap)-len(m.maxHeap) > 1 {
		heap.Push(&m.maxHeap, heap.Pop(&m.minHeap))
	}

}

// Median returns the median of the numbers
func (m *Median) Median() float64 {
	minHeapLen := m.minHeap.Len()
	maxHeapLen := m.maxHeap.Len()

	if maxHeapLen == 0 && minHeapLen == 0 {
		return 0
	} else if maxHeapLen < minHeapLen {
		return float64(m.minHeap[0])
	}

	return float64(m.maxHeap[0])
}

func load(s string) []int {
	inputBytes, err := ioutil.ReadFile(s)

	poof(err)

	inputString := string(inputBytes[:])
	rows := strings.Split(inputString, "\n")

	size := len(rows) - 1

	data := make([]int, size)

	for j := 0; j < size; j++ {
		is := rows[j]
		i, err := strconv.Atoi(strings.TrimSpace(is))

		poof(err)

		data[j] = i
	}

	return data
}

func poof(err error) {
	if err != nil {
		panic(err)
	}
}
