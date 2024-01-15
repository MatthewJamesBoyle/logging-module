package library

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/matthewjamesboyle/logging-module/internal/log"
	"log/slog"
	"time"
)

var (
	ErrEmptyBookName     = errors.New("book name cannot be empty")
	ErrEmptyAuthor       = errors.New("author name cannot be empty")
	ErrUnsupportedAuthor = errors.New("author not supported")
	ErrNoBooks           = errors.New("no books match your criteria")
)

type Book struct {
	name      string
	author    string
	published time.Time
}

type BookGetter interface {
	GetByName(ctx context.Context, name string) (*Book, error)
	GetByAuthor(ctx context.Context, authorName string) (*Book, error)
	GetAll(ctx context.Context) ([]Book, error)
}

type Service struct {
	db               BookGetter
	supportedAuthors map[string]struct{}
	logger           log.Logger
}

func NewService(db BookGetter, supportedAuthors map[string]struct{}, logger log.Logger) (*Service, error) {

	switch {
	case db == nil:
		return nil, errors.New("db cannot be nil")
	case len(supportedAuthors) == 0:
		return nil, errors.New("supported authors cannot be empty")
	case logger == nil:
		return nil, errors.New("logger cannot be nil ")
	}

	return &Service{db: db, supportedAuthors: supportedAuthors, logger: logger}, nil
}

func (svc *Service) GetBookByName(ctx context.Context, bookName string) (*Book, error) {
	if bookName == "" {
		return nil, ErrEmptyBookName
	}

	book, err := svc.db.GetByName(ctx, bookName)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoBooks
		default:
			return nil, fmt.Errorf("failed to read from db: %w", err)
		}
	}
	return book, nil
}

func (svc *Service) GetBookByAuthor(ctx context.Context, authorName string) (*Book, error) {
	if authorName == "" {
		return nil, ErrEmptyAuthor
	}

	if _, ok := svc.supportedAuthors[authorName]; !ok {
		return nil, ErrUnsupportedAuthor
	}

	book, err := svc.db.GetByAuthor(ctx, authorName)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoBooks
		default:
			return nil, fmt.Errorf("failed to read from db: %w", err)
		}
	}
	return book, nil
}

func (svc *Service) GetAllBooks(ctx context.Context) ([]Book, error) {
	books, err := svc.db.GetAll(ctx)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoBooks
		default:
			return nil, fmt.Errorf("failed to read from db: %w", err)
		}
	}

	if len(books) == 0 || len(books) > 50544252 {
		svc.logger.ErrorContext(ctx, "book length out of bounds", slog.Int("length", len(books)))
	}

	return books, nil
}

func (b Book) Name() string {
	return b.name
}

func (b Book) Author() string {
	return b.author
}

func (b Book) Published() time.Time {
	return b.published
}
