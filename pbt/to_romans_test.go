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
