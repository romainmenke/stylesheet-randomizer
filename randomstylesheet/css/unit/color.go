package unit

import (
	"fmt"
)

type Color struct {
	r Value
	g Value
	b Value
	a Value

	withAlpha bool
}

func (x Color) Css() string {
	if x.withAlpha {
		return fmt.Sprintf("rgba(%s,%s,%s,%s)", x.r.Css(), x.g.Css(), x.b.Css(), x.a.Css())
	}

	return fmt.Sprintf("rgb(%s,%s,%s)", x.r.Css(), x.g.Css(), x.b.Css())
}

func (x Color) UnitType() string {
	return "color"
}

func (x Color) Unit() string {
	if x.withAlpha {
		return "rgba"
	}

	return "rgb"
}

func (x Color) Value() interface{} {
	if x.withAlpha {
		return struct {
			R interface{}
			G interface{}
			B interface{}
			A interface{}
		}{
			R: x.r.Value(),
			G: x.g.Value(),
			B: x.b.Value(),
			A: x.a.Value(),
		}
	}

	return struct {
		R interface{}
		G interface{}
		B interface{}
	}{
		R: x.r.Value(),
		G: x.g.Value(),
		B: x.b.Value(),
	}
}

func ColorRGB() Value {
	return Color{
		r: NewInteger(Bit8 | Absolute),
		g: NewInteger(Bit8 | Absolute),
		b: NewInteger(Bit8 | Absolute),
	}
}

func ColorRGBA() Value {
	return Color{
		r: NewInteger(Bit8 | Absolute),
		g: NewInteger(Bit8 | Absolute),
		b: NewInteger(Bit8 | Absolute),
		a: NewNumber(Zero | Decimals | Absolute),
	}
}
