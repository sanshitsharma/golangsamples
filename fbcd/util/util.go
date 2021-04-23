package util

import (
	"errors"
	"fmt"
)

/******************************************************************************
//    FBCD: Facebook Compressor-Decompressor. A modified LZSS implementation
//
//    File    : util.go
//    Purpose : Declare all constants and variables which are used across
//              multiple packages
//    Author  : Sanshit Sharma(sanshit.sharma@gmail.com)
//    Date    : Since April 2021
******************************************************************************/

/********************************************************************************************************/

const (
	// When a repeatation is found, we will use 16bits to encode the offset.
	offsetBits = 16
	// Using 16 bits allows out max offset to be 2^16 - 1 = 65535, which gives
	// us the size of out search buffer
	SearchBufferSize = 1 << offsetBits

	// Upon finding a repeatition, we can use upto 6 bits to specify the copy
	// length.
	lengthBits = 6
	// Additionally, as stated in the problem, there is not use compressing
	// matches of length 2 or less as encoding a 2 length match will take 23
	// bits whereas writing them out unencoded will take 18bits. Therefore,
	// we can define a max uncompressed length
	MaxUncompressed = 2
	// Using the 6 bits for copy length and 2 length of uncompressed gives
	// us the size of the look ahead buffer
	LookAheadBufferSize = (1 << lengthBits) + MaxUncompressed
)

var (
	ErrEmptyBuffer = errors.New(`buffer is empty`)
)

type Queue struct {
	buffer []byte
	front  int32
	rear   int32
	size   uint32
}

func initQ(size uint32) *Queue {
	return &Queue{
		buffer: make([]byte, size),
		front: -1,
		rear: -1,
		size: size,
	}
}

func (q *Queue) Enqueue(datum byte) error {
	// If queue is full, instead of rejecting, we will just drop
	// a byte from the head
	if IncIdx(q.rear, int32(q.size)) == q.front {
		// dequeue 1 element
		_, err := q.Dequeue()
		if err != nil && err != ErrEmptyBuffer {
			return err
		}
	}

	// empty queue
	if q.front == -1 {
		q.front = 0
		q.rear = 0
	} else {
		q.rear = IncIdx(q.rear, int32(q.size))
	}

	q.buffer[q.rear] = datum
	return nil
}

func (q *Queue) Dequeue() (byte, error) {
	// If queue is empty return nil
	if q.front == -1 {
		return byte(0), ErrEmptyBuffer
	}

	var res byte
	// Only 1 element in queue
	if q.front == q.rear {
		res = q.buffer[q.front]
		q.front = -1
		q.rear = -1
	} else {
		res = q.buffer[q.front]
		q.front = IncIdx(q.front, int32(q.size))
	}

	return res, nil
}

// Peek will return the element at front of the queue without actually removing it
func (q *Queue) Peek() (byte, error) {
	// If queue is empty return nil
	if q.front == -1 {
		return byte(0), ErrEmptyBuffer
	}

	return q.buffer[q.front], nil
}

func (q *Queue) GetFront() int32 {
	return q.front
}

func (q *Queue) GetRear() int32 {
	return q.rear
}

func (q *Queue) GetSize() uint32 {
	return q.size
}

func (q *Queue) GetBuffer() []byte {
	return q.buffer
}

// IsEmpty return if queue is empty
func (q *Queue) IsEmpty() bool {
	return q.front == -1
}

// GetCount returns the current count of elements in the buffer
func (q *Queue) GetCount() uint32 {
	if q.front == -1 {
		return 0
	}
	if q.front > q.rear {
		return q.size - uint32(q.front) + uint32(q.rear) + 1
	}

	return uint32(q.rear - q.front + 1)
}

// GetDist returns the number of elements between two indexes of the buffer
func (q *Queue) GetDist(idx1, idx2 int32) uint32 {
	if idx1 > idx2 {
		return q.size - uint32(idx1) + uint32(idx2) + 1
	}

	return uint32(idx2 - idx1 + 1)
}

func (q *Queue) String() string {
	return fmt.Sprintf("Value: %v, Front: %v, Rear: %v, Size: %v", string(q.buffer), q.front, q.rear, q.size)
}

func (q *Queue) Stats() string {
	return fmt.Sprintf("Size: %v, Front: %v, Rear: %v, Contents: %v", q.size, q.front, q.rear, q.buffer)
}

// IncIdx returns the value of a particular index cyclically incremented by 1
func IncIdx(idx, size int32) int32 {
	return (idx+1)%size
}

// SlidingWindow struct stores the search buffer and the uncompressed
// look ahead buffers
type SlidingWindow struct {
	Search    *Queue
	LookAhead *Queue
}

func (sw *SlidingWindow) String() string {
	return fmt.Sprintf("Search: '%v'. Look Ahead: '%v'", string(sw.Search.buffer), string(sw.LookAhead.buffer))
}

func (sw *SlidingWindow) Stats() string {
	return fmt.Sprintf("Search Buffer. '%v'\nLookAhead Buffer. '%v'", sw.Search.Stats(), sw.LookAhead.Stats())
}

func InitSlidingWindow() *SlidingWindow {
	return &SlidingWindow{
		Search:    initQ(SearchBufferSize),
		LookAhead: initQ(LookAheadBufferSize),
	}
}

/********************************************************************************************************/

/********************************************************************************************************/
/*                                              HELPERS                                                 */
/********************************************************************************************************/

// UnrotateIndex returns the unrotated index value for a slice
func UnrotateIndex(value, limit uint32) uint32 {
	if value < limit {
		return value
	}

	return value - limit
}
