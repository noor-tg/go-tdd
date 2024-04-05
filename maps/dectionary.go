package maps

import (
	"errors"
)

type Dectionary map[string]string

var NotFoundOnDectionary = errors.New("word Not Found on dectionary")

func (d Dectionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", NotFoundOnDectionary
	}
	return definition, nil
}
