package unit

import (
	"fmt"
)

type Integer int

func (x Integer) Css() string {
	v := int(x)
	return fmt.Sprint(v)
}

func (x Integer) UnitType() string {
	return "integer"
}

func (x Integer) Unit() string {
	return "integer"
}

func (x Integer) Value() interface{} {
	return int(x)
}

func NewInteger(scale Scale) Value {
	value := scale.RndValue()
	if floatValue, ok := value.(float64); ok {
		return Integer(int(floatValue))
	}

	if intValue, ok := value.(int); ok {
		return Integer(intValue)
	}

	// fallback to '1' if random scaling fails
	return Integer(1)
}
