package naive

import (
	"github.com/sanshitsharma/golangsamples/fbcd/matcher"
	"github.com/sanshitsharma/golangsamples/fbcd/util"
)

// Matcher is the underlying type which implements the Matcher interface
// used to find the longest matching pattern
type Matcher struct{}

func Create() *Matcher {
	return &Matcher{}
}

// GetLongest returns the longest matching pattern between the Search and LookAhead buffer
// This is a brute force implementation which take O(mn) time where n is the length of the
// search buffer and m is the length of the look ahead buffer
func (m *Matcher) GetLongest(sw *util.SlidingWindow) (*matcher.Encoded, error) {
	longestMatch := 0
	startSearch := -1
	startLookAhead := -1
	//fmt.Println("***********************************************")
	//fmt.Printf("finding longest match!!\nSearch<head: %v, tail:%v, numElems: %v>, Look Ahead<head: %v, tail:%v, numElems: %v>\n",
	//	sw.Search.GetFront(), sw.Search.GetRear(), sw.Search.GetCount(),
	//	sw.LookAhead.GetFront(), sw.LookAhead.GetRear(), sw.LookAhead.GetCount())

	//fmt.Println("Search isEmpty:", sw.Search.IsEmpty(), " Look Ahead isEmpty:", sw.LookAhead.IsEmpty())

	if !sw.Search.IsEmpty() && !sw.LookAhead.IsEmpty() {
		for i := int(sw.Search.GetFront()); i != int(sw.Search.GetRear()); i = (i + 1) % int(sw.Search.GetSize()) {
			counter := 0
			j := int(sw.LookAhead.GetFront())
			for sw.Search.GetBuffer()[i+counter] == sw.LookAhead.GetBuffer()[j+counter] {
				counter++
				if ((i + counter) >= int(sw.Search.GetSize())) || ((j + counter) >= int(sw.LookAhead.GetSize())) {
					//fmt.Println("breaking!!!")
					break
				} else if counter > longestMatch {
					longestMatch = counter
					startSearch = i
					startLookAhead = j
				}
			}
		}
	}
	//fmt.Println("found longest match =", longestMatch)
	//fmt.Println("***********************************************")

	// If the longest match is less than 3, then just ignore it
	if longestMatch <= util.MaxUncompressed {
		datum, err := sw.LookAhead.Peek()
		if err != nil {
			return nil, err
		}
		return &matcher.Encoded{
			Type:   matcher.UNCOMPRESSED,
			Datum:  datum,
			Offset: 0,
			Length: 0,
		}, nil
	}

	//fmt.Printf("longest match found!! search start idx: %v, look ahead start idx: %v\n",
	//	startSearch, startLookAhead)
	// If not, then calculate the offset and return
	// offset = dist of startSearch from search.rear + dist of startLookAhead from lookahead.front
	offset := sw.Search.GetDist(int32(startSearch), sw.Search.GetRear()) +
		sw.LookAhead.GetDist(sw.LookAhead.GetFront(), int32(startLookAhead)) - 1

	return &matcher.Encoded{
		Type:   matcher.COMPRESSED,
		Offset: offset,
		Length: uint8(longestMatch),
	}, nil
}

// ShiftBytes moves n bytes from the beginning of the LookAhead buffer
// to the tail of the SearchBuffer
func (m *Matcher) ShiftBytes(sw *util.SlidingWindow, n uint8) error {
	if sw == nil {
		return matcher.ErrEmptyBuffer
	}

	if n == 0 {
		// nothing to read
		return nil
	}

	// fmt.Println("Before shifting..", sw.LookAhead.Stats())

	for i := 0; i < int(n) && i < int(sw.LookAhead.GetSize()); i++ {
		datum, err := sw.LookAhead.Dequeue()
		if err != nil {
			return err
		}

		err = sw.Search.Enqueue(datum)
		if err != nil {
			return err
		}
	}

	return nil
}
