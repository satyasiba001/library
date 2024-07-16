package services

import (
	"context"
	"fmt"

	"github.com/satyasiba001/library/repository"
)

type LibraryHandler struct {
	db repository.LibraryRepository
}

func NewLibraryHandler(db repository.LibraryRepository) *LibraryHandler {
	return &LibraryHandler{
		db: db,
	}
}

func (lh *LibraryHandler) Ping(ctx context.Context) error {
	errr := lh.db.Ping(ctx)
	if errr != nil {
		fmt.Println("22222", errr)
		return errr
	}
	return nil
}
