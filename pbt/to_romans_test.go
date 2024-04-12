package pbt

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	arabic int
	roman  string
}{
	{
		arabic: 1,
		roman:  "I",
	},
	{
		arabic: 2,
		roman:  "II",
	},
	{
		arabic: 3,
		roman:  "III",
	},
	{
		arabic: 4,
		roman:  "IV",
	},
	{
		arabic: 5,
		roman:  "V",
	},
	{
		arabic: 6,
		roman:  "VI",
	},
	{
		arabic: 7,
		roman:  "VII",
	},
	{
		arabic: 8,
		roman:  "VIII",
	},
	{
		arabic: 9,
		roman:  "IX",
	},
	{
		arabic: 10,
		roman:  "X",
	},
	{
		arabic: 11,
		roman:  "XI",
	},
	{
		arabic: 12,
		roman:  "XII",
	},
	{
		arabic: 13,
		roman:  "XIII",
	},
	{
		arabic: 14,
		roman:  "XIV",
	},
	{
		arabic: 15,
		roman:  "XV",
	},
	{
		arabic: 16,
		roman:  "XVI",
	},
	{
		arabic: 17,
		roman:  "XVII",
	},
	{
		arabic: 18,
		roman:  "XVIII",
	},
	{
		arabic: 19,
		roman:  "XIX",
	},
	{
		arabic: 20,
		roman:  "XX",
	},
	{
		arabic: 29,
		roman:  "XXIX",
	},
	{
		arabic: 30,
		roman:  "XXX",
	},
	{
		arabic: 39,
		roman:  "XXXIX",
	},
	{
		arabic: 47,
		roman:  "XLVII",
	},
	{
		arabic: 40,
		roman:  "XL",
	},
	{
		arabic: 50,
		roman:  "L",
	},
	{
		arabic: 99,
		roman:  "XCIX",
	},
	{
		arabic: 100,
		roman:  "C",
	},
	{
		arabic: 129,
		roman:  "CXXIX",
	},
	{
		arabic: 150,
		roman:  "CL",
	},
	{
		arabic: 498,
		roman:  "CDXCVIII",
	},
	{
		arabic: 500,
		roman:  "D",
	},
	{
		arabic: 1984,
		roman:  "MCMLXXXIV",
	},
	{
		arabic: 1993,
		roman:  "MCMXCIII",
	},
	{
		arabic: 4680,
		roman:  "MMMMDCLXXX",
	},
	{
		arabic: 10000,
		roman:  "MMMMMMMMMM",
	},
}

func TestRomanNumerals(t *testing.T) {
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%v converted to %v\n", tC.arabic, tC.roman), func(t *testing.T) {
			got := ConvertToRoman(tC.arabic)

			if got != tC.roman {
				t.Errorf("got %v want %v", got, tC.roman)
			}
		})
	}
}

func TestToArabic(t *testing.T) {
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%v convert to %v", tC.roman, tC.arabic), func(t *testing.T) {
			got := ConvertToArabic(tC.roman)

			if got != tC.arabic {
				t.Errorf("got %v, want %v", got, tC.arabic)
			}
		})
	}

}

func ExampleRomanNumberFifty() {
	arabic := 50
	got := ConvertToRoman(arabic)

	fmt.Printf("arabic = %d , roman = %q", arabic, got)
	//Output:
	// arabic = 50 , roman = "L"
}
func ExampleRomanNumberFortyNine() {
	arabic := 49
	got := ConvertToRoman(arabic)

	fmt.Printf("arabic = %d , roman = %q", arabic, got)
	//Output:
	// arabic = 49 , roman = "XLIX"
}
