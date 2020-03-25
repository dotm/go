package main

func main() {}

//Dictionary represents real world dictionary
type Dictionary map[string]string

const (
	//ErrNotFound is used when a dictionary doesn't have the entry for a word
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	//ErrWordExists is used when we try adding a key that already exist in a dictionary
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	//ErrWordDoesNotExist is used when we try updating a key that is nonexistent in a dictionary
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

//DictionaryErr is a wrapper so that error for dictionaries can be defined as constant
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

//Add inserts a new word into a dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

//Update modify a word's definition if it already exists
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

//Search finds a word in a dictionary by key
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

//Delete a word in a dictionary whether it exist or not
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
