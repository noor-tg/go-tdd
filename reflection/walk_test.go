package reflection_test

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name:          "struct with one string field",
			Input:         struct{ Name string }{"Alnoor"},
			ExpectedCalls: []string{"Alnoor"},
		},
		{
			Name: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Alnoor", "Cairo"},
			ExpectedCalls: []string{"Alnoor", "Cairo"},
		},
		{
			Name: "struct with two string field one int field",
			Input: struct {
				Name string
				Age  int
			}{"Alnoor", 10},
			ExpectedCalls: []string{"Alnoor"},
		},
	}

	for _, test := range cases {

		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %q, expected %q", got, test.ExpectedCalls)
			}

		})
	}
}

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}

}
