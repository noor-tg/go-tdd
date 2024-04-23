package generics

import "errors"

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}
func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	length := len(s.values)
	value := s.values[length-1]
	s.values = s.values[:length-1]
	return value, nil
}
