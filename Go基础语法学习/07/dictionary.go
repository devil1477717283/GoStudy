package _7

func Search(dictionary map[string]string, dest string) string {
	return dictionary[dest]
}

type Dictionary map[string]string

const (
	NotFound = DictionaryErr("Not Found Word in Dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
func (d Dictionary) Search(word string) (string, error) {
	Destination, ok := d[word]
	if !ok {
		return Destination, NotFound
	}
	return Destination, nil
}
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
