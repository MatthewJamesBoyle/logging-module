package transport

import (
	"encoding/json"
	"errors"
	"github.com/matthewjamesboyle/logging-module/internal/library"
	"github.com/matthewjamesboyle/logging-module/internal/log"
	"log/slog"
	"net/http"
)

type BookResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description,omitempty"`
	PublishedOn string `json:"published_on"`
	Genre       string `json:"genre,omitempty"`
}

type Handler struct {
	svc    library.Service
	logger log.Logger
}

func NewHandler(svc library.Service, logger log.Logger) (*Handler, error) {
	return &Handler{svc: svc, logger: logger}, nil
}

func (h Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	reqID := slog.String("request_id", r.Context().Value(requestIDKey{}).(string))

	books, err := h.svc.GetAllBooks(ctx)
	if err != nil {
		switch {
		case errors.Is(err, library.ErrEmptyBookName):
			h.logger.InfoContext(
				r.Context(),
				"empty_book_passed",
				reqID,
			)
			http.Error(w, "Book name is required", http.StatusBadRequest)
		case errors.Is(err, library.ErrNoBooks):
			h.logger.InfoContext(
				r.Context(),
				"no_books_found",
				reqID,
			)
			http.Error(w, "No book found with given name", http.StatusNotFound)
		default:
			h.logger.ErrorContext(
				r.Context(),
				"internal_server_error",
				slog.Any("err", err),
				reqID,
			)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	var bookRes = make([]*BookResponse, 0)
	for _, v := range books {
		bookRes = append(bookRes, newBookResponse(&v))
	}

	response, err := json.Marshal(bookRes)
	if err != nil {
		h.logger.ErrorContext(
			r.Context(),
			"failed_to_marshal_response",
			slog.Any("err", err),
			reqID,
		)
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		h.logger.ErrorContext(
			r.Context(),
			"failed_to_write_response",
			slog.Any("err", err),
			reqID,
		)
	}
}

func newBookResponse(book *library.Book) *BookResponse {
	return &BookResponse{
		Title:       book.Name(),
		Author:      book.Author(),
		PublishedOn: book.Published().Format("2006-01-02"),
	}
}
