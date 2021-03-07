package maps

const (
	ErrNotFound          = DictErr("word not found")
	ErrWordExists        = DictErr("word existed")
	ErrWordDoesNotExists = DictErr("word existed")
)

type DictErr string

func (e DictErr) Error() string {
	return string(e)
}

type Dict map[string]string

func (dict Dict) Search(word string) (string, error) {
	value, ok := dict[word]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dict) Add(word, value string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = value
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dict) Update(word, value string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = value
		return nil
	case ErrNotFound:
		return ErrWordDoesNotExists
	default:
		return err
	}
}

func (d Dict) Delete(word string) {
	delete(d, word)
}
