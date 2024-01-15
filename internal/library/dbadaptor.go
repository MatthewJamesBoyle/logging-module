package library

import (
	"context"
	"github.com/matthewjamesboyle/logging-module/internal/db"
)

type MockAdaptor struct {
	db db.MockDb
}

func NewMockAdaptor(db db.MockDb) *MockAdaptor {
	return &MockAdaptor{db: db}
}

func (m MockAdaptor) GetByName(ctx context.Context, name string) (*Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockAdaptor) GetByAuthor(ctx context.Context, authorName string) (*Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockAdaptor) GetAll(ctx context.Context) ([]Book, error) {
	books, err := m.db.GetAllBooks(ctx)
	if err != nil {
		return nil, err
	}

	var tbooks = make([]Book, len(books))
	for i := range books {
		tbooks[i] = Book{
			name:      books[i].Title,
			author:    books[i].Author,
			published: books[i].PublishedOn,
		}
	}

	return tbooks, nil
}
