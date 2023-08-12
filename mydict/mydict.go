package mydict

import "errors"

var (
	errNotFound   = errors.New("Not found")
	errCantUpdate = errors.New("Cant update non-existing word")
	errWordExists = errors.New("Word already exists")
)

// MyDict type
type MyDict map[string]string

// Search for a word
func (d MyDict) Search(word string) (string, error) {
	if value, exists := d[word]; exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d MyDict) Add(word, def string) error {
	switch _, err := d.Search(word); err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word
func (d MyDict) Update(word, def string) error {
	switch _, err := d.Search(word); err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d MyDict) Delete(word string) {
	delete(d, word)
}
