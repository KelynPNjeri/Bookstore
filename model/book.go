package model

import (
	"github.com/google/uuid"
	"time"
)

type Book struct {
	ID uuid.UUID `json:"id"`
	BookName string `json:"book_name"`
	BookDesc string `json:"book_desc"`
	Category string `json:"category"`
	DateAdded time.Time
}

type Books []Book
