package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFind      = errors.New("could not find the word you are looking for")
	ErrExistingWord = errors.New("it is not possible to add the word because it already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, exists := d[word]
	if !exists {
		return "", ErrNotFind
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFind:
		d[word] = definition
	case nil:
		return ErrExistingWord
	default:
		return err
	}
	return nil
}
