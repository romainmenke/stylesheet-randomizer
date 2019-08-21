package unit

import (
	"math"
	mr "math/rand"
)

type Scale uint64

const (
	Absolute          Scale = 1 << iota // only positive numbers
	Decimals                            // with decimals numbers
	Zero                                // 0
	One                                 // 1
	Binary                              // 0-1
	Small                               // 0-3
	Medium                              // 3-8
	Large                               // 8-13
	XLarge                              // 13-34
	XXLarge                             // 34-55
	XXXLarge                            // 55-144
	Bit8                                // 0-255
	Percentage                          // 0-100
	PercentagePlus100                   // 100-200
)

func (x Scale) Set(flag Scale) Scale {
	return x | flag
}

func (x Scale) Clear(flag Scale) Scale {
	return x &^ flag
}

func (x Scale) Toggle(flag Scale) Scale {
	return x ^ flag
}

func (x Scale) Has(flag Scale) bool {
	return x&flag != 0
}

func (x Scale) RndValue() interface{} {
	decimalsNumbers := x.Has(Decimals)
	absoluteNumbers := x.Has(Absolute)

	selected := []Scale{}
	for _, flag := range []Scale{Zero, One, Binary, Small, Medium, Large, XLarge, XXLarge, XXXLarge, Bit8, Percentage, PercentagePlus100} {
		if x.Has(flag) {
			selected = append(selected, flag)
		}
	}

	if len(selected) == 0 {
		return 0
	}

	chosen := selected[mr.Intn(len(selected))]

	wholePart := 0

	switch chosen {
	case Zero: // 0
		wholePart = 0
	case One: // 1
		wholePart = 1
	case Binary: // 0-1
		wholePart = mr.Intn(1 + 1)
	case Small: // 0-3
		wholePart = mr.Intn(3 + 1)
	case Medium: // 3-8
		wholePart = mr.Intn(5+1) + 3
	case Large: // 8-13
		wholePart = mr.Intn(5+1) + 8
	case XLarge: // 13-34
		wholePart = mr.Intn(21+1) + 13
	case XXLarge: // 34-55
		wholePart = mr.Intn(21+1) + 34
	case XXXLarge: // 55-144
		wholePart = mr.Intn(89+1) + 55

	case Bit8: // 0-255
		wholePart = mr.Intn(255 + 1)

	case Percentage: // 0-100
		wholePart = mr.Intn(100 + 1)
	case PercentagePlus100: // 100-200
		wholePart = mr.Intn(100+1) + 100
	}

	// whole positive number
	if !decimalsNumbers && absoluteNumbers {
		return wholePart
	}

	if !decimalsNumbers && !absoluteNumbers {
		if mr.Intn(1) == 0 {
			// whole positive number
			return wholePart
		}

		// whole negative number
		return wholePart * -1
	}

	decimalPart := mr.ExpFloat64()
	decimalPart = math.Abs(decimalPart - math.Round(decimalPart))
	number := decimalPart + float64(wholePart)
	if absoluteNumbers {
		// decimal positive number
		return number
	}

	if mr.Intn(1) == 0 {
		// decimal positive number
		return number
	}

	// decimal negative number
	return number * -1
}
