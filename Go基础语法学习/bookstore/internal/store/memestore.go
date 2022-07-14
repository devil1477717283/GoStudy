package store

import (
	mystore "bookstore/store"
	factory "bookstore/store/factory"
	"fmt"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}

func (m *MemStore) Create(book *mystore.Book) error {
	if _, ok := m.books[book.Id]; ok {
		return fmt.Errorf("Book already exsited!")
	}
	m.books[book.Id] = book
	return nil
}
func (m *MemStore) Update(book *mystore.Book) error {
	if _, ok := m.books[book.Id]; !ok {
		return fmt.Errorf("Book don't exsited!")
	}
	m.books[book.Id] = book
	return nil
}
func (m *MemStore) Get(Id string) (mystore.Book, error) {
	if _, ok := m.books[Id]; !ok {
		return mystore.Book{}, fmt.Errorf("Book don't exsited!")
	}
	return *m.books[Id], nil
}
func (m *MemStore) Delete(Id string) error {
	if _, ok := m.books[Id]; ok {
		delete(m.books, Id)
	} else {
		return fmt.Errorf("Book don't exsited!")
	}
	return nil
}
func (m MemStore) GetAll() []mystore.Book {
	books := make([]mystore.Book, 0)
	for _, book := range m.books {
		books = append(books, *book)
	}
	return books
}
