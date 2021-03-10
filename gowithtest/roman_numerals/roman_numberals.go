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

var allRomanNumeralsDict = func() map[string]int {
	dict := make(map[string]int)
	for _, v := range allRomanNumerals {
		dict[v.Symbol] = v.Value
	}
	return dict
}()

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	return allRomanNumeralsDict[string(symbols)]
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
	for _, symbols := range parseSymbols(roman) {
		total += allRomanNumerals.ValueOf(symbols...)
	}
	return total
}

func parseSymbols(roman string) (symbols [][]byte) {
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]
		isNotAtEnd := (i+1 < len(roman))

		if isNotAtEnd && isSubtractive(symbol) && allRomanNumerals.ValueOf(symbol, roman[i+1]) != 0 {
			symbols = append(symbols, []byte{symbol, roman[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return symbols
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
