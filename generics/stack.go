package generics

import "errors"

type StackOfInts struct {
	values []int
}

func (s *StackOfInts) IsEmpty() bool {
	return len(s.values) == 0
}
func (s *StackOfInts) Push(value int) {
	s.values = append(s.values, value)
}

func (s *StackOfInts) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	length := len(s.values)
	value := s.values[length-1]
	s.values = s.values[:length-1]
	return value, nil
}

type StackOfStrings struct {
	values []string
}

func (s *StackOfStrings) IsEmpty() bool {
	return len(s.values) == 0
}
func (s *StackOfStrings) Push(value string) {
	s.values = append(s.values, value)
}

func (s *StackOfStrings) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("stack is empty")
	}
	length := len(s.values)
	value := s.values[length-1]
	s.values = s.values[:length-1]
	return value, nil
}
