package integers

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

func TestAdder(t *testing.T) {

	sum := Add(1, 2)
	expected := 3

	testza.AssertEqual(t, sum, expected)

}
