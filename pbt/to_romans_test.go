package pbt

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

var testCases = []struct {
	arabic ConvertibleArabic
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
}

func TestRomanNumerals(t *testing.T) {
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%v converted to %v\n", tC.arabic, tC.roman), func(t *testing.T) {
			got, err := ConvertToRoman(tC.arabic)

			if err != nil {
				t.Error(err)
			} else if got != tC.roman {
				t.Errorf("got %v want %v", got, tC.roman)
			}
		})
	}
}

func TestToArabic(t *testing.T) {
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%v convert to %v", tC.roman, tC.arabic), func(t *testing.T) {
			got, err := ConvertToArabic(tC.roman)

			if err != nil {
				t.Error(err)
			} else if got != tC.arabic {
				t.Errorf("got %v, want %v", got, tC.arabic)
			}
		})
	}

}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic ConvertibleArabic) bool {
		if arabic < 0 || arabic > 3999 {
			log.Println(arabic)
			return true
		}
		roman, _ := ConvertToRoman(arabic)
		fromRoman, _ := ConvertToArabic(roman)

		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("assertion failed", err)
	}

}

func ExampleRomanNumberFifty() {
	var arabic ConvertibleArabic = 50
	got, _ := ConvertToRoman(arabic)

	fmt.Printf("arabic = %d , roman = %q", arabic, got)
	//Output:
	// arabic = 50 , roman = "L"
}
func ExampleRomanNumberFortyNine() {
	var arabic ConvertibleArabic = 49
	got, _ := ConvertToRoman(arabic)

	fmt.Printf("arabic = %d , roman = %q", arabic, got)
	//Output:
	// arabic = 49 , roman = "XLIX"
}
