package unit

type Value interface {
	Css() string
	UnitType() string
	Unit() string
	Value() interface{}
}
