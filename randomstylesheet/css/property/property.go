package property

import (
	"encoding/json"
	"fmt"

	"github.com/romainmenke/stylesheet-randomizer/randomstylesheet/css/unit"
)

type Rule struct {
	Property string
	Value    unit.Value
}

func (x Rule) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Property string      `json:"property"`
		UnitType string      `json:"unit_type"`
		Unit     string      `json:"unit"`
		Value    interface{} `json:"value"`
	}{
		Property: x.Property,
		UnitType: x.Value.UnitType(),
		Unit:     x.Value.Unit(),
		Value:    x.Value.Value(),
	})
}

func (x Rule) Css() string {
	return fmt.Sprintf("%s: %s;", x.Property, x.Value.Css())
}

func List() []func() Rule {
	return []func() Rule{
		BackgroundColor,
		Color,
		LineHeight,
		PaddingLeft,
		PaddingTop,
		PaddingBottom,
		PaddingRight,
	}
}

func BackgroundColor() Rule {
	return Rule{
		Property: "background-color",
		Value: unit.NewSet(
			unit.ColorRGB,
			unit.ColorRGBA,
			unit.Transparant,
			unit.Inherit,
			unit.Initial,
			unit.Unset,
		),
	}
}

func Color() Rule {
	return Rule{
		Property: "color",
		Value: unit.NewSet(
			unit.ColorRGB,
			unit.ColorRGBA,
			unit.Transparant,
			unit.Inherit,
			unit.Initial,
			unit.Unset,
		),
	}
}

func LineHeight() Rule {
	return Rule{
		Property: "line-height",
		Value: unit.NewSet(
			unit.Unitless,
			unit.LengthEm,
			unit.LengthRem,
			unit.LengthPercentage,
			unit.Normal,
			unit.Inherit,
			unit.Initial,
			unit.Unset,
		),
	}
}

func FontSize() Rule {
	return Rule{
		Property: "font-size",
		Value: unit.NewSet(
			unit.Unitless,
			unit.LengthEm,
			unit.LengthRem,
			unit.LengthPercentage,
			unit.Normal,
			unit.Inherit,
			unit.Initial,
			unit.Unset,
		),
	}
}

func PaddingLeft() Rule {
	return Rule{
		Property: "padding-left",
		Value: unit.NewSet(
			unit.Unitless,
			unit.LengthEm,
			unit.LengthRem,
			unit.LengthPercentage,
			unit.Normal,
			unit.Inherit,
			unit.Initial,
			unit.Unset,
		),
	}
}

func PaddingTop() Rule {
	return Rule{
		Property: "padding-top",
		Value: unit.NewSet(
			unit.Unitless,
			unit.LengthEm,
			unit.LengthRem,
			unit.LengthPercentage,
			unit.Normal,
			unit.Inherit,
			unit.Initial,
			unit.Unset,
		),
	}
}

func PaddingBottom() Rule {
	return Rule{
		Property: "padding-bottom",
		Value: unit.NewSet(
			unit.Unitless,
			unit.LengthEm,
			unit.LengthRem,
			unit.LengthPercentage,
			unit.Normal,
			unit.Inherit,
			unit.Initial,
			unit.Unset,
		),
	}
}

func PaddingRight() Rule {
	return Rule{
		Property: "padding-right",
		Value: unit.NewSet(
			unit.Unitless,
			unit.LengthEm,
			unit.LengthRem,
			unit.LengthPercentage,
			unit.Normal,
			unit.Inherit,
			unit.Initial,
			unit.Unset,
		),
	}
}
