package roman_numerals

import (
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func (r RomanNumerals) ValueOf(symbol string) int {
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

// ConvertToRoman return roman string number from arabic number
func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

// ConvertToNumeric return arabic from roman string number
func ConvertToNumeric(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		// look ahead
		if i+1 < len(roman) && isSubtractive(symbol) {
			nextSymbol := roman[i+1]

			potentialGroup := string([]byte{symbol, nextSymbol})

			if value := allRomanNumerals.ValueOf(potentialGroup); value != 0 {
				total += value
				i++
			} else {
				total += allRomanNumerals.ValueOf(string([]byte{symbol}))
			}
		} else {
			total += allRomanNumerals.ValueOf(string([]byte{symbol}))
		}
	}
	return total
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
