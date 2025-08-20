package maps

import "errors"

type Dictionary map[string]string

var ErrNotFind = errors.New("could not find the word you are looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, exists := d[word]
	if !exists {
		return "", ErrNotFind
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
