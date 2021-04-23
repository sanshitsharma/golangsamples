package matcher

import (
	"errors"
	"fmt"
	"github.com/sanshitsharma/golangsamples/fbcd/util"
)

var (
	ErrEmptyBuffer = errors.New("uninitialized buffer")
)

// Matcher interface aggregates all the methods for a performing longest pattern
// matching between two buffers
type Matcher interface {
	// GetLongest returns the longest match
	GetLongest(sw *util.SlidingWindow) (*Encoded, error)

	// ReplaceByte replaces the byte at a given index with a new byte
	//ReplaceByte(sw *util.SlidingWindow, idx uint32, replacement byte)

	// ShiftBytes moves n bytes from the beginning of the LookAhead buffer
	// to the tail of the SearchBuffer
	ShiftBytes(sw *util.SlidingWindow, n uint8) uint8
}

/********************************************************************************************************/

type Encoding uint8

const (
	UNCOMPRESSED Encoding = iota
	COMPRESSED
)

// String returns the string value of Encoding
func (encoding Encoding) String() string {
	return string(encoding)
}

// Encoded struct stores compressed bytes in (offset, length) format.
// This struct will be parsed to before writing the actual compressed
// bit output to the file
type Encoded struct {
	Type   Encoding
	// fields populated when Type == UNCOMPRESSED
	Datum  byte
	// fields populated when Type == COMPRESSED
	Offset uint32
	Length uint8
}

func (enc *Encoded) String() string {
	if enc.Type == UNCOMPRESSED {
		return fmt.Sprintf("-----------> 0'%v' <-----------", string(enc.Datum))
	}

	return fmt.Sprintf("-----------> (1, %v, %v) <-----------", enc.Offset, enc.Length)
}

/********************************************************************************************************/