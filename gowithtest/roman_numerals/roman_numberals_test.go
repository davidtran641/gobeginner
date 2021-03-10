package roman_numerals

import (
	"fmt"
	"testing"
)

var cases = []struct {
	arabic int
	roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{90, "XC"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, tt := range cases {
		tc := tt
		t.Run(fmt.Sprintf("Covert %d", tc.arabic), func(t *testing.T) {
			got := ConvertToRoman(tc.arabic)
			want := tc.roman
			if got != want {
				t.Errorf("want %q but got %q", want, got)
			}
		})
	}
}

func TestConvertToNumeric(t *testing.T) {
	for _, tt := range cases {
		tc := tt
		t.Run(fmt.Sprintf("Covert %s to %d", tc.roman, tc.arabic), func(t *testing.T) {
			got := ConvertToNumeric(tc.roman)
			want := tc.arabic
			if got != want {
				t.Errorf("want %d but got %d", want, got)
			}
		})
	}
}
