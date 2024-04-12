package pbt

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

// roman numbers based on symbol
var AllRomanNumerals = []RomanNumeral{
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

func ConvertToRoman(number int) string {
	var result strings.Builder
	// iterate throw roman numbers
	for _, numeral := range AllRomanNumerals {
		// if my current number is equal or larger than this number
		for number >= numeral.Value {
			// write the symbol
			result.WriteString(numeral.Symbol)
			// subtract the value
			number -= numeral.Value
		}
	}
	return result.String()
}

func ConvertToArabic(roman string) int {
	arabic := 0
	for _, numeral := range AllRomanNumerals {
		// if my current number is equal or larger than this number
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return arabic
}
