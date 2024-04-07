package dep_inject

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Noor")

	got := buffer.String()
	want := "Hello, Noor"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
