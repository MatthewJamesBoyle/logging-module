package db

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"time"
)

type Book struct {
	ID          string    `db:"id"`           // Unique identifier for the book
	Title       string    `db:"title"`        // Title of the book
	Author      string    `db:"author"`       // Author of the book
	Description string    `db:"description"`  // A short description of the book
	PublishedOn time.Time `db:"published_on"` // Publication date of the book
	Genre       string    `db:"genre"`        // Genre of the book
}

type MockDb struct{}

func (m MockDb) GetAllBooks(ctx context.Context) ([]Book, error) {
	randomNumber := rand.Intn(3) + 1

	switch randomNumber {
	case 1:
		return []Book{
			{
				ID:          "1",
				Title:       "The Great Adventure",
				Author:      "Jane Doe",
				Description: "An exciting journey through uncharted territories.",
				PublishedOn: time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC),
				Genre:       "Adventure",
			},
			{
				ID:          "2",
				Title:       "Mystery of the Lost City",
				Author:      "John Smith",
				Description: "A thrilling mystery set in a forgotten city.",
				PublishedOn: time.Date(2018, 5, 23, 0, 0, 0, 0, time.UTC),
				Genre:       "Mystery",
			},
			{
				ID:          "3",
				Title:       "Science and You",
				Author:      "Alice Johnson",
				Description: "Exploring the wonders of science in everyday life.",
				PublishedOn: time.Date(2021, 8, 15, 0, 0, 0, 0, time.UTC),
				Genre:       "Science",
			},
		}, nil
	case 2:
		return nil, sql.ErrNoRows
	case 3:
		return nil, errors.New("some unknown error")
	}
	return nil, nil
}
