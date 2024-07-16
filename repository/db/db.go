package db

import (
	"context"
	"database/sql"
	"fmt"

	// "errors"
	_ "github.com/lib/pq"
	"github.com/satyasiba001/library/repository"
)

type LibraryDb struct {
	Db *sql.DB
}

func NewLibraryRepository(db *sql.DB) repository.LibraryRepository {
	return LibraryDb{
		Db: db,
	}
}

func (dbo LibraryDb) Ping(ctx context.Context) error {
	err := dbo.Db.Ping()
	if err != nil {
		fmt.Println("11111",err)
		return err
	}
	return nil
}
