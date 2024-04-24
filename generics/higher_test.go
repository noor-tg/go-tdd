package generics

import (
	"strings"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("sum two functions by reduce", func(t *testing.T) {
		got := Sum([]int{1, 2, 3})
		want := 6

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

type Person struct {
	Name string
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		want := 2
		num, found := Find(numbers, func(number int) bool {
			return number%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, num, want)
	})
	t.Run("find last name in list of full names", func(t *testing.T) {
		people := []Person{
			{
				Name: "ali",
			},
			{
				Name: "osama",
			},
			{
				Name: "ahmed",
			},
			{
				Name: "alnoor",
			},
		}
		person, found := Find(people, func(person Person) bool {
			return strings.Contains(person.Name, "alnoor")
		})
		AssertTrue(t, found)
		AssertEqual(t, person, Person{Name: "alnoor"})
	})
}
