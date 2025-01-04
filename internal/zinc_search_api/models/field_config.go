package zinc_search

import "regexp"

type FieldConfig struct {
	Regex      *regexp.Regexp
	AddValue   func(matches []string)
	SplitValue bool
}
