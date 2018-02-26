package heap

import (
	"database/sql/driver"
	"errors"
)

// Type is an interger type which represents the types of heaps
type Type int

const (
	// None represents an invalid heap type
	None Type = 0
	// Min Type represents a min heap
	Min Type = 1
	// Max Type represents a max heap
	Max Type = 2
)

// Value returns the integer value of the Type heaptype enum
func (heapType Type) Value() (driver.Value, error) {
	switch heapType {
	case Min:
		break
	case Max:
		break
	default:
		return nil, errors.New("invalid heaptype")
	}
	return driver.Value(int(heapType)), nil
}

// Scan converts a numreic value to a heap type enum
func (heapType *Type) Scan(value interface{}) {
	switch value.(type) {
	case int:
		*heapType = Type(value.(int))
		break
	default:
		*heapType = Type(None)
	}
}

// ToString converts heaptype to string
func (heapType Type) toString() string {
	switch heapType {
	case Min:
		return "Min"
	case Max:
		return "Max"
	default:
		return ""
	}
}
