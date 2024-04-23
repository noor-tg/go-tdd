package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Grace")
	})
}

func AssertNotEqual[T comparable](t *testing.T, i1, i2 T) {
	t.Helper()
	if i1 == i2 {
		t.Errorf("i1 %+v, i2 %+v", i1, i2)
	}
}

func AssertEqual[T comparable](t *testing.T, i1, i2 T) {
	t.Helper()
	if i1 != i2 {
		t.Errorf("i1 %+v, i2 %+v", i1, i2)
	}
}
func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v", got)
	}
}
