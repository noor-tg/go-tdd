package pbt

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	testCases := []struct {
		input int
		want  string
	}{
		{
			input: 1,
			want:  "I",
		},
		{
			input: 2,
			want:  "II",
		},
		{
			input: 3,
			want:  "III",
		},
		{
			input: 4,
			want:  "IV",
		},
		{
			input: 5,
			want:  "V",
		},
		{
			input: 6,
			want:  "VI",
		},
		{
			input: 7,
			want:  "VII",
		},
		{
			input: 8,
			want:  "VIII",
		},
		{
			input: 9,
			want:  "IX",
		},
		{
			input: 10,
			want:  "X",
		},
		{
			input: 11,
			want:  "XI",
		},
		{
			input: 12,
			want:  "XII",
		},
		{
			input: 13,
			want:  "XIII",
		},
		{
			input: 14,
			want:  "XIV",
		},
		{
			input: 15,
			want:  "XV",
		},
		{
			input: 16,
			want:  "XVI",
		},
		{
			input: 17,
			want:  "XVII",
		},
		{
			input: 18,
			want:  "XVIII",
		},
		{
			input: 19,
			want:  "XIX",
		},
		{
			input: 20,
			want:  "XX",
		},
		{
			input: 29,
			want:  "XXIX",
		},
		{
			input: 30,
			want:  "XXX",
		},
		{
			input: 39,
			want:  "XXXIX",
		},
		{
			input: 47,
			want:  "XLVII",
		},
		{
			input: 40,
			want:  "XL",
		},
		{
			input: 50,
			want:  "L",
		},
		{
			input: 99,
			want:  "XCIX",
		},
		{
			input: 100,
			want:  "C",
		},
		{
			input: 129,
			want:  "CXXIX",
		},
		{
			input: 150,
			want:  "CL",
		},
		{
			input: 498,
			want:  "CDXCVIII",
		},
		{
			input: 500,
			want:  "D",
		},
		{
			input: 1984,
			want:  "MCMLXXXIV",
		},
		{
			input: 1993,
			want:  "MCMXCIII",
		},
		{
			input: 4680,
			want:  "MMMMDCLXXX",
		},
		{
			input: 10000,
			want:  "MMMMMMMMMM",
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%d converted to %q\n", tC.input, tC.want), func(t *testing.T) {

			got := ConvertToRoman(tC.input)

			if got != tC.want {
				t.Errorf("got %q want %q", got, tC.want)
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
