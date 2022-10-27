package utils

import (
	"fmt"
	"strings"
)

// Utility for search specific item in list
func Contains(list []interface{}, data interface{}) bool {
	for _, v := range list {
		if v == data {
			return true
		}
	}
	return false
}

/**
Params
@raw: Number in string format.
@unit: Currency.
@precision: Number of decimal places of currency.
@separator: Regex that separates the number in thousand.
@decimalSep: Regex to separate the @precision.
@format: Return the conversion result.
*/
func CurrencyConverter(raw, unit, separator, decimalSep, format string, precision int) string {
	runes := []rune(raw)

	var temp []string
	if len(raw) > 3 {
		var raws_ []string
		var newstr string

		for i := len(raw) - 1; i >= 0; i-- {
			raws_ = append(raws_, fmt.Sprintf("%c", runes[i]))

			if len(raws_)%3 != 0 {
				newstr = strings.Join(raws_, "")
			}

			if len(raws_)%3 == 0 {
				newstr = strings.Join(raws_, "")
				temp = append(temp, newstr)
				newstr = ""
				raws_ = []string{}
			}
		}

		var final string
		if len(raws_) > 0 {
			temp = append(temp, newstr)
			final = strings.Join(temp, separator)
		} else {
			final = strings.Join(temp, separator)
		}

		var rvs = ""
		final_runes := []rune(final)
		for i := len(final) - 1; i >= 0; i-- {
			rvs = rvs + fmt.Sprintf("%c", final_runes[i])
		}

		return fmt.Sprintf(format+CurrencyWithPrecision(decimalSep, precision), unit, rvs)
	}

	return fmt.Sprintf(format+CurrencyWithPrecision(decimalSep, precision), unit, raw)

}

func CurrencyWithPrecision(decimalSep string, precision int) string {
	var init = "0"

	if precision < 0 {
		precision = 2
	}

	newstr := strings.Repeat(init, precision)
	newstr = decimalSep + newstr
	return newstr

}
