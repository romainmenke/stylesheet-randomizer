package selector

import "github.com/romainmenke/stylesheet-randomizer/randomstylesheet/css/property"

var selectors = []string{
	"h1",
	"h2",
	"h3",
	"h4",
	"h5",
	"h6",
	"ul",
	"ol",
	"li",
	"hr",
	"p",
	"table",
	"thead",
	"tbody",
	"tfoot",
	"tr",
	"th",
	"td",
	"blockquote",
	"cite",
	"pre",
	"code",
	"a",
	"strong",
	"em",
	"i",
	"s",
}

func List() []string {
	return selectors
}

type Selector struct {
	Name  string          `json:"name"`
	Rules []property.Rule `json:"rules"`
}

const newLine = `
`

func (x Selector) Css() string {
	out := x.Name + " {" + newLine

	for _, rule := range x.Rules {
		out = out + rule.Css() + newLine
	}

	out = out + "}" + newLine + newLine

	return out
}

type StyleSheet []Selector

func (x StyleSheet) Css() string {
	out := ""

	for _, selector := range x {
		out = out + selector.Css()
	}

	return out
}
