package pbt

import "testing"

func TestRomanNumerals(t *testing.T) {
	testCases := []struct {
		desc  string
		input int
		want  string
	}{
		{
			desc:  "1 converted to I",
			input: 1,
			want:  "I",
		},
		{
			desc:  "2 converted to II",
			input: 2,
			want:  "II",
		},
		{
			desc:  "3 converted to III",
			input: 3,
			want:  "III",
		},
		{
			desc:  "4 converted to IV",
			input: 4,
			want:  "IV",
		},
		{
			desc:  "5 converted to V",
			input: 5,
			want:  "V",
		},
		{
			desc:  "6 converted to VI",
			input: 6,
			want:  "VI",
		},
		{
			desc:  "7 converted to VII",
			input: 7,
			want:  "VII",
		},
		{
			desc:  "8 converted to VIII",
			input: 8,
			want:  "VIII",
		},
		{
			desc:  "9 converted to IX",
			input: 9,
			want:  "IX",
		},
		{
			desc:  "10 converted to X",
			input: 10,
			want:  "X",
		},
		{
			desc:  "11 converted to XI",
			input: 11,
			want:  "XI",
		},
		{
			desc:  "12 converted to XII",
			input: 12,
			want:  "XII",
		},
		{
			desc:  "13 converted to XIII",
			input: 13,
			want:  "XIII",
		},
		{
			desc:  "14 converted to XIV",
			input: 14,
			want:  "XIV",
		},
		{
			desc:  "15 converted to XV",
			input: 15,
			want:  "XV",
		},
		{
			desc:  "16 converted to XVI",
			input: 16,
			want:  "XVI",
		},
		{
			desc:  "17 converted to XVII",
			input: 17,
			want:  "XVII",
		},
		{
			desc:  "18 converted to XVIII",
			input: 18,
			want:  "XVIII",
		},
		{
			desc:  "19 converted to XIX",
			input: 19,
			want:  "XIX",
		},
		{
			desc:  "20 converted to XX",
			input: 20,
			want:  "XX",
		},
		{
			desc:  "29 converted to XXIX",
			input: 29,
			want:  "XXIX",
		},
		{
			desc:  "30 converted to XXX",
			input: 30,
			want:  "XXX",
		},
		{
			desc:  "39 converted to XXXIX",
			input: 39,
			want:  "XXXIX",
		},
		{
			desc:  "40 converted to XXXX",
			input: 40,
			want:  "XXXX",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			got := ConvertToRoman(tC.input)

			if got != tC.want {
				t.Errorf("got %q want %q", got, tC.want)
			}
		})
	}
}
