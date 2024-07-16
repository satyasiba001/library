package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/satyasiba001/library/libhttp"
	"github.com/satyasiba001/library/repository/db"
	"github.com/satyasiba001/library/services"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "library"
	sslmode  = "disabl"
)

func main() {
	psqlconn := fmt.Sprintf("host= %s port = %d user = %s password= %s dbname= %s sslmode= %s", host, port, user, password, dbname, sslmode)

	// connection open
	sqlDb, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	libraryRepository := db.NewLibraryRepository(sqlDb)

	libraryHandler := services.NewLibraryHandler(libraryRepository)
	engine := gin.Default()

	library:= libhttp.NewLibraryHttp(engine,libraryHandler)

	library.Routes()
	errr := engine.Run(":8000")
	if errr!= nil {
		fmt.Println("app not started")
	}
	
}
