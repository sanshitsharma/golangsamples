package main

import (
	"fmt"
	"github.com/dlintw/goconf"
	"strings"
)

var (
	compBitPosition    map[string]uint16
	compIdx            map[uint16]string
	categoryMasks      map[string]uint16
)

func main() {
	fmt.Println("Hello app!!")

	err := loadDependencies()

	fmt.Println("Categories initialized:\n", "red_mask", categoryMasks["Red"], "yellow_mask", categoryMasks["Yellow"],
		"comp_bit_pos", compBitPosition, "comp_idx", compIdx, "err", err)
}

func loadDependencies() error {
	var shift uint16

	compBitPosition = make(map[string]uint16)
	compIdx = make(map[uint16]string)
	categoryMasks = make(map[string]uint16)

	// Load red comps
	lenRedComps, err := initCategory("Red", shift)
	if err != nil {
		return err
	}
	fmt.Println("Length Red Comps:", lenRedComps)

	// Load yellow state components
	shift += uint16(lenRedComps)
	lenYellowComps, err := initCategory("Yellow", shift)
	fmt.Println("Length Yellow Comps:", lenYellowComps)

	return err
}

func initCategory(category string, shift uint16) (int, error) {
	comps, err := readConf(category)
	if err != nil {
		return 0, err
	}

	var cMask uint16

	// Calculate the bit position of the components in the target status
	for indx, comp := range comps {
		compBitPosition[comp] = shift + uint16(indx)
		compIdx[shift + uint16(indx)] = comp
		cMask |= 1 << uint16(indx)
	}

	// Shift and store the category mask
	cMask <<= shift
	categoryMasks[category] = cMask

	return len(comps), nil
}

func readConf(category string) ([]string, error) {
	c, err := goconf.ReadConfigFile("dependencyLoader/app.conf")
	if err != nil {
		fmt.Println("failed to read conf file:", err)
		return nil, err
	}

	comps, err := c.GetString(`manager`, category)
	if err != nil {
		return nil, err
	}

	return strings.Split(comps, `,`), nil
}
