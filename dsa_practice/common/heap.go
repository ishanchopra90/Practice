package common

type MinHeap[T any] struct {
	Heap    []T
	Compare func(a, b T) int
}

func (minHeap *MinHeap[T]) Add(heapElement T) {
	minHeap.Heap = append(minHeap.Heap, heapElement)

	// minheapify upwards
	i := len(minHeap.Heap) - 1
	for i >= 0 {
		var parentI int
		if i%2 == 0 {
			parentI = (i - 2) / 2
		} else {
			parentI = (i - 1) / 2
		}

		if parentI < 0 {
			break
		}

		if minHeap.Compare(minHeap.Heap[i], minHeap.Heap[parentI]) == -1 {
			temp := minHeap.Heap[parentI]
			minHeap.Heap[parentI] = minHeap.Heap[i]
			minHeap.Heap[i] = temp
			i = parentI
		} else {
			break
		}
	}
}

func (minHeap *MinHeap[T]) RemoveMin() T {
	// swap first and last element, remove last element, minHeapify downwards
	elementToReturn := minHeap.Heap[0]
	temp := minHeap.Heap[0]
	minHeap.Heap[0] = minHeap.Heap[len(minHeap.Heap)-1]
	minHeap.Heap[len(minHeap.Heap)-1] = temp
	minHeap.Heap = minHeap.Heap[:len(minHeap.Heap)-1]
	i := 0
	for i < len(minHeap.Heap) {
		c1 := 2*i + 1
		c2 := 2*i + 2
		if c1 >= len(minHeap.Heap) && c2 >= len(minHeap.Heap) {
			break
		}

		smallerChild := 0
		if c1 >= len(minHeap.Heap) {
			smallerChild = c2
		} else if c2 >= len(minHeap.Heap) {
			smallerChild = c1
		} else if minHeap.Compare(minHeap.Heap[c1], minHeap.Heap[c2]) == -1 {
			smallerChild = c1
		} else {
			smallerChild = c2
		}

		if minHeap.Compare(minHeap.Heap[smallerChild], minHeap.Heap[i]) == -1 {
			temp := minHeap.Heap[smallerChild]
			minHeap.Heap[smallerChild] = minHeap.Heap[i]
			minHeap.Heap[i] = temp
			i = smallerChild
		} else {
			break
		}
	}

	return elementToReturn
}
