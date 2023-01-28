package main

import (
	"container/heap"
	"fmt"
	"strings"
	"time"
)

type intHeap []int

// Len is the number of elements in the collection.
func (h intHeap) Len() int {
	return len(h)
}

// Less reports whether the element with index i
// must sort before the element with index j.
// This will determine whether the heap is a min heap or a max heap
func (h intHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

// Swap swaps the elements with indexes i and j.
func (h intHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Each holiday has a name and a date
type Holiday struct {
	name string
	date time.Time
}

// for convenience, we can create a string representation to print the holiday instance
func (h Holiday) String() string {
	return "[" + h.name + ", " + h.date.Format("Jan 2") + "]"
}

// We can create a new type to represent a list of holidays
type Holidays []Holiday

// Implement the container/heap interface
func (h Holidays) Len() int { return len(h) }
func (h Holidays) Less(i, j int) bool {
	return strings.Compare(h[i].name, h[j].name) < 0
}
func (h Holidays) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Holidays) Push(x interface{}) {
	*h = append(*h, x.(Holiday))
}

func (h *Holidays) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	nums := &intHeap{3, 1, 4, 5, 1, 1, 2, 6}

	heap.Init(nums)
	fmt.Println(nums)

	min := heap.Pop(nums)
	fmt.Println("min: ", min, " heap: ", *nums)

	heap.Push(nums, 5)
	fmt.Println("heap: ", *nums)

	sortedNums := []int{}
	for nums.Len() > 0 {
		sortedNums = append(sortedNums, heap.Pop(nums).(int))
	}
	fmt.Println("sorted values: ", sortedNums)

	holidays := &Holidays{}
	heap.Init(holidays)

	heap.Push(holidays, Holiday{name: "Christmas", date: time.Date(2023, time.December, 25, 0, 0, 0, 0, time.Local)})
	heap.Push(holidays, Holiday{name: "Labour Day", date: time.Date(2023, time.May, 1, 0, 0, 0, 0, time.Local)})
	heap.Push(holidays, Holiday{name: "Diwali", date: time.Date(2023, time.November, 23, 0, 0, 0, 0, time.Local)})

	fmt.Println("holidays: ", holidays)
}
