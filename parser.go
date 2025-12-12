package unitor

import (
	"regexp"
	"strconv"
	"strings"
)

type NumberParser interface {
	Parse(input string, unitSuffix string) (value float64, unit string, err error)
}

var DefaultParser NumberParser = &BasicNumberParser{}

type BasicNumberParser struct {
}

var (
	numberMatcher = regexp.MustCompile(`[\d][\d.,]*`)
	unitMatcher   = regexp.MustCompile(`[^\s\d.,]+`)
)

func (BasicNumberParser) Parse(input string, unitSymbol string) (value float64, unit string, err error) {
	// Find unit
	potentialUnits := unitMatcher.FindAllString(input, 3)
	if len(potentialUnits) > 0 {
		if unitSymbol == "" {
			// No help given, just use last match as most units have it in back
			unit = potentialUnits[len(potentialUnits)-1]
		} else {
			for _, pu := range potentialUnits {
				// Take first potential unit that contains unit symbol
				if strings.Contains(pu, unitSymbol) {
					unit = pu
					break
				}
			}
		}
	}

	// Find number
	potentialValues := numberMatcher.FindAllString(input, 2)
	if len(potentialValues) > 0 {
		// Take first
		rawValue := potentialValues[0]

		// Parse
		value, err = strconv.ParseFloat(rawValue, 64)
	}
	return
}
