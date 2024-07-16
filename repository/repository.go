package repository

import "context"

type LibraryRepository interface {
	Ping(context.Context) error
}
