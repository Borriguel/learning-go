package maps

type Dictionary map[string]string
type ErrDictionary string

const (
	ErrNotFind         = ErrDictionary("could not find the word you are looking for")
	ErrExistingWord    = ErrDictionary("it is not possible to add the word because it already exists")
	ErrNonExistingWord = ErrDictionary("it is not possible to update the word because it does not exists")
)

func (e ErrDictionary) Error() string {
	return string(e)
}

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

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFind:
		return ErrNonExistingWord
	case nil:
		d[word] = newDefinition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
