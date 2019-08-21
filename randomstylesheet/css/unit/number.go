package unit

import (
	"fmt"
)

type Number float64

func (x Number) Css() string {
	v := float64(x)
	return fmt.Sprintf("%.4f", v)
}

func (x Number) UnitType() string {
	return "number"
}

func (x Number) Unit() string {
	return "number"
}

func (x Number) Value() interface{} {
	return float64(x)
}

func NewNumber(scale Scale) Value {
	value := scale.RndValue()
	if floatValue, ok := value.(float64); ok {
		return Number(floatValue)
	}

	if intValue, ok := value.(int); ok {
		return Number(float64(intValue))
	}

	// fallback to '1' if random scaling fails
	return Number(1)
}

func Unitless() Value {
	return NewNumber(
		Zero |
			One |
			Small |
			Medium |
			Large |
			XLarge |
			XXLarge |
			XXXLarge |
			Decimals |
			Absolute,
	)
}
