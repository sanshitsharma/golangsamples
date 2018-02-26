package heap

import (
	"errors"
	"fmt"
)

// Heap struct stores the heap elements and the type of the heap
type Heap struct {
	elems []int
	hType Type
}

// NewHeap returns a new empty heap
func NewHeap(heapType Type) (*Heap, error) {
	if _, err := heapType.Value(); err != nil {
		return nil, errors.New("invalid heap type")
	}
	return &Heap{
		elems: make([]int, 0),
		hType: heapType,
	}, nil
}

func (heap *Heap) comparator(childIndx int, parentIndx int) bool {
	if (heap.hType == Max && heap.elems[childIndx] > heap.elems[parentIndx]) ||
		(heap.hType == Min && heap.elems[childIndx] < heap.elems[parentIndx]) {
		return true
	}

	return false
}

// Bottom up, make the heap
func (heap *Heap) heapify() {
	cI := len(heap.elems) - 1
	pI := (cI - 1) / 2

	for pI >= 0 && heap.comparator(cI, pI) {
		// swap the elements
		swap(heap.elems, cI, pI)

		cI = pI
		pI = (cI - 1) / 2
	}
}

// Insert an element to the heap and heapify according to the type of the heap
func (heap *Heap) Insert(value int) {
	heap.elems = append(heap.elems, value)
	heap.heapify()
}

func swap(elems []int, i, j int) {
	temp := elems[i]
	elems[i] = elems[j]
	elems[j] = temp
}

// Top down heapify staring at index indx
func (heap *Heap) fix(indx int) bool {
	// Check how many children does this node have
	c1I := 2*indx + 1
	c2I := 2*indx + 2

	numChildren := 0
	if c1I >= 0 && c1I <= len(heap.elems)-1 && c2I >= 0 && c2I <= len(heap.elems)-1 {
		numChildren = 2
	} else if c1I >= 0 && c1I <= len(heap.elems)-1 {
		numChildren = 1
	}

	// recursion termination condition
	if numChildren == 0 || (numChildren == 1 && !heap.comparator(c1I, indx)) || (numChildren == 2 && !heap.comparator(c1I, indx) && !heap.comparator(c2I, indx)) {
		return true
	}

	// target node has 1 child
	if numChildren == 1 && heap.comparator(c1I, indx) {
		// swap
		swap(heap.elems, indx, c1I)
		return heap.fix(c1I)
	}

	//target node has two children, then swap with the appropriate child
	if heap.elems[c1I] > heap.elems[c2I] && heap.comparator(c1I, indx) {
		swap(heap.elems, indx, c1I)
		return heap.fix(c1I)
	} else if heap.elems[c1I] < heap.elems[c2I] && heap.comparator(c2I, indx) {
		swap(heap.elems, indx, c2I)
		return heap.fix(c2I)

	}

	return false
}

// Delete an element from heap if it exists
func (heap *Heap) Delete(value int) error {
	found := false
	for indx, elem := range heap.elems {
		if elem == value {
			//fmt.Printf("found element [%v] at indx [%v]\n", elem, indx)
			found = true
			// Copy element at last index at target index
			heap.elems[indx] = heap.elems[len(heap.elems)-1]
			// Delete the last element
			heap.elems = heap.elems[:len(heap.elems)-1]

			// Heapify top down starting at target index
			heap.fix(indx)
			break
		}
	}

	if !found {
		return fmt.Errorf("element [%v] does not exist in heap", value)
	}

	return nil
}

// Print the heap
func (heap *Heap) Print() {
	fmt.Println("Heap Type:", heap.hType.toString())
	fmt.Println("Elements:", heap.elems)
	//fmt.Println(heap.elems)
}

// Size returns the number of elements in the heap
func (heap *Heap) Size() int {
	return len(heap.elems)
}

// Pop returns the top element of the heap and deletes that element from the heap
func (heap *Heap) Pop() (int, error) {
	if heap.IsEmpty() {
		return -1, errors.New(`heap is empty`)
	}

	elem := heap.elems[0]
	heap.Delete(elem)
	return elem, nil
}

// Peek returns the top element of the heap but doesn't delete it
func (heap *Heap) Peek() (int, error) {
	if heap.IsEmpty() {
		return -1, errors.New(`heap is empty`)
	}

	return heap.elems[0], nil
}

func (heap *Heap) IsEmpty() bool {
	return len(heap.elems) == 0
}
