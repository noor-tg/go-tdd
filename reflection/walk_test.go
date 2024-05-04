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
	// function which get int return reflect value
	var getField func(int) reflect.Value

	// switch throw value kind
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	// can use multiple types with switch case by ,
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		// set function
		getField = val.Index
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
		//maps need special handling because it only get by key and it is not ordered
	case reflect.Map:
		// get map keys
		for _, key := range val.MapKeys() {
			// get value by key
			walk(val.MapIndex(key).Interface(), fn)
		}
	// channels need custom fetch
	case reflect.Chan:
		// loop infinitely
		for {
			// for each value received if it ok wall throw it
			if v, ok := val.Recv(); ok {
				walk(v.Interface(), fn)
			} else {
				// break out of loop if not receive any
				break
			}
		}
	// handle function type
	case reflect.Func:
		// call function with nil
		result := val.Call(nil)
		// iterate throw result values
		for _, res := range result {
			walk(res.Interface(), fn)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		// use numberof value + getfield method to walk throw values
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	// get value of interface
	val := reflect.ValueOf(x)

	// check if kind of value is pointer
	if val.Kind() == reflect.Pointer {
		// return value of pointer
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
