package pbt

import (
	"errors"
	"strings"
)

type RomanNumeral struct {
	Value  ConvertibleArabic
	Symbol string
}

type ConvertibleArabic uint16

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

func ConvertToRoman(number ConvertibleArabic) (string, error) {
	if number > 3999 {
		return "I", errors.New("can not convert number larger than 3999")
	}
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
	return result.String(), nil
}

func ConvertToArabic(roman string) (ConvertibleArabic, error) {
	var arabic ConvertibleArabic
	for _, numeral := range AllRomanNumerals {
		// if my current number is equal or larger than this number
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	if arabic > 3999 {
		return 1, errors.New("converting number larger than 3999 is in-correct")
	}

	return arabic, nil
}
