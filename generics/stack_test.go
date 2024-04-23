package generics

import "testing"

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		stack := new(Stack[int])

		AssertTrue(t, stack.IsEmpty())

		stack.Push(1)
		AssertFalse(t, stack.IsEmpty())

		stack.Push(2)

		value, _ := stack.Pop()
		AssertEqual(t, value, 2)
		AssertFalse(t, stack.IsEmpty())

		value2, _ := stack.Pop()
		AssertEqual(t, value2, 1)
		AssertTrue(t, stack.IsEmpty())
	})

	t.Run("strings stack", func(t *testing.T) {
		stack := new(Stack[string])

		AssertTrue(t, stack.IsEmpty())

		stack.Push("a")
		AssertFalse(t, stack.IsEmpty())

		stack.Push("b")

		value, _ := stack.Pop()
		AssertEqual(t, value, "b")
		AssertFalse(t, stack.IsEmpty())

		value2, _ := stack.Pop()
		AssertEqual(t, value2, "a")
		AssertTrue(t, stack.IsEmpty())
	})

}
