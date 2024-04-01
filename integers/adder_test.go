package integers

import (
	"fmt"
	"testing"

	"github.com/MarvinJWendt/testza"
)

func TestAdder(t *testing.T) {

	sum := Add(1, 2)
	expected := 3

	testza.AssertEqual(t, sum, expected)

}

func ExampleAdd() {
	sum := Add(2, 2)
	fmt.Println(sum)
	// Output: 4
}

func init() {
	testza.SetShowStartupMessage(false)
}
