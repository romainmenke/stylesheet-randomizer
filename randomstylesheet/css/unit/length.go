package unit

import (
	"fmt"
)

type Length struct {
	value Value
	unit  string
}

func (x Length) Css() string {
	return fmt.Sprintf("%s%s", x.value.Css(), x.unit)
}

func (x Length) UnitType() string {
	return "length"
}

func (x Length) Unit() string {
	return x.unit
}

func (x Length) Value() interface{} {
	return x.value.Value()
}

func LengthPixels() Value {
	return Length{
		unit: "px",
		value: NewNumber(
			Small |
				Medium |
				Large |
				XLarge |
				XXLarge |
				XXXLarge,
		),
	}
}

func LengthEm() Value {
	return Length{
		unit: "em",
		value: NewNumber(
			Zero |
				One |
				Small |
				Decimals |
				Absolute,
		),
	}
}

func LengthRem() Value {
	return Length{
		unit: "rem",
		value: NewNumber(
			Zero |
				One |
				Small |
				Decimals |
				Absolute,
		),
	}
}

func LengthPercentage() Value {
	return Length{
		unit: "%",
		value: NewInteger(
			Percentage |
				PercentagePlus100 |
				Absolute,
		),
	}
}
