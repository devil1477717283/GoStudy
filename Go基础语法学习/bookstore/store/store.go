package store

type Book struct {
	Id      string   `json:"id"`
	name    string   `json:"name"`
	Authors []string `json:"authors"'`
	Press   string   `json:"press"`
}

type Store interface {
	Create(*Book) error
	Update(*Book) error
	Get(string) (Book, error)
	Delete(string) error
	GetAll() []Book
}
