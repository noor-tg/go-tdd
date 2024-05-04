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
		{
			Name:          "arrays contain strings",
			Input:         [2]Profile{{"Cairo", 10}, {"Khartoum", 5}},
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
	t.Run("walk with maps", func(t *testing.T) {
		var got []string

		input := map[string]string{"city": "Cairo", "kind": "male"}

		walk(input, func(input string) {
			got = append(got, input)
		})

		assertContains(t, "Cairo", got)
		assertContains(t, "male", got)
	})
	t.Run("walk with channels", func(t *testing.T) {
		chnl := make(chan Profile)

		go func() {
			chnl <- Profile{"Khartoum", 10}
			chnl <- Profile{"Cairo", 5}
			close(chnl)
		}()
		var got []string
		want := []string{"Khartoum", "Cairo"}

		walk(chnl, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, expected %v", got, want)
		}
	})
	t.Run("walk with functions", func(t *testing.T) {
		afunc := func() (Profile, Profile) {
			return Profile{"Khartoum", 10}, Profile{"Cairo", 5}
		}
		var got []string
		want := []string{"Khartoum", "Cairo"}

		walk(afunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, expected %v", got, want)
		}
	})
}

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walk(v.Interface(), fn)
			} else {
				break
			}
		}
	case reflect.Func:
		result := val.Call(nil)
		for _, res := range result {
			walk(res.Interface(), fn)
		}
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

func assertContains(t testing.TB, needle string, heystacke []string) {
	t.Helper()
	contains := false

	for _, val := range heystacke {
		if needle == val {
			contains = true
		}
	}

	if !contains {
		t.Errorf("map %+v not contain %q", heystacke, needle)
	}
}
