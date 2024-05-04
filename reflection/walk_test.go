package reflection_test

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	type Profile struct {
		City string
		Age  int
	}
	type Person struct {
		Name    string
		Profile Profile
	}

	type Developer struct {
		Person
		Skills []string
	}

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
		{
			Name:          "struct with nested structs",
			Input:         Person{"Alnoor", Profile{"Cairo", 10}},
			ExpectedCalls: []string{"Alnoor", "Cairo"},
		}, {
			Name:          "struct with nested structs as pointer",
			Input:         &Person{"Alnoor", Profile{"Cairo", 10}},
			ExpectedCalls: []string{"Alnoor", "Cairo"},
		},
		{
			Name:          "slices contain strings",
			Input:         []Profile{{"Cairo", 10}, {"Khartoum", 5}},
			ExpectedCalls: []string{"Cairo", "Khartoum"},
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
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.Slice:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.String:
		fn(val.String())
		return
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
