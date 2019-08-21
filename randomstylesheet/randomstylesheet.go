package randomstylesheet

import (
	"bytes"
	"encoding/json"
	"io"
	mr "math/rand"

	"github.com/romainmenke/stylesheet-randomizer/randomstylesheet/css/property"
	"github.com/romainmenke/stylesheet-randomizer/randomstylesheet/css/selector"
)

// Write a random styleheet to the provided writer
func Write(w io.Writer, seed int64) {
	w.Write([]byte(NewStylesheet(seed).Css()))
}

func JSON(seed int64, pretty bool) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)

	if pretty {
		enc.SetIndent("", "    ")
	}

	err := enc.Encode(struct {
		Seed       int64 `json:"seed"`
		StyleSheet selector.StyleSheet
	}{
		Seed:       seed,
		StyleSheet: NewStylesheet(seed),
	})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func NewStylesheet(seed int64) selector.StyleSheet {
	mr.Seed(seed)

	sheet := []selector.Selector{}

	for _, selectorName := range selector.List() {
		sheet = append(sheet, selector.Selector{
			Name:  selectorName,
			Rules: NewRules(),
		})
	}

	return selector.StyleSheet(sheet)
}

func NewRules() []property.Rule {
	properties := property.List()

	numberOfRules := mr.Intn(len(properties))
	if numberOfRules == 0 {
		return []property.Rule{}
	}

	rulesUnique := map[int]struct{}{}
	ruleIndices := []int{}
	for i := 0; i < numberOfRules; i++ {
		idx := mr.Intn(len(properties))
		if _, ok := rulesUnique[idx]; ok {
			continue
		}

		rulesUnique[idx] = struct{}{}
		ruleIndices = append(ruleIndices, idx)
	}

	rules := []property.Rule{}

	for index := range ruleIndices {
		rules = append(rules, properties[index]())
	}

	return rules
}
