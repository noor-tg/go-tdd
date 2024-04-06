package maps

type Dectionary map[string]string
type DectionaryErr string

// implement error interface by making Error method
func (e DectionaryErr) Error() string {
	return string(e)
}

const (
	NotFoundOnDectionary = DectionaryErr("word Not Found in dectionary")
	WordAlreadyExist     = DectionaryErr("word already exist in dectionary")
)

func (d Dectionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", NotFoundOnDectionary
	}
	return definition, nil
}

func (d Dectionary) Add(word string, definition string) error {
	// search for word in dect
	// return err if exist
	_, err := d.Search(word)
	// check err value
	switch err {
	// if not found add word def
	case NotFoundOnDectionary:
		d[word] = definition
	// if err nil means word already exist
	// so throw error
	case nil:
		return WordAlreadyExist
	// if not eather case return the err
	default:
		return err
	}
	// no error if code execute here
	return nil
}
