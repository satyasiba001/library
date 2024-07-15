package main

import (
	"fmt"
	"database/sql"
	// "errors"
	_ "github.com/lib/pq"
)

const(
		host   		= "localhost"
		port   		= 5432
		user   		= "postgres"
	password		= "postgres"
	dbname			= "library"
	sslmode 		= "disable"
)

func dbConnection()(*sql.DB, error){
	// connection string
	psqlconn := fmt.Sprintf("host= %s port = %d user = %s password= %s dbname= %s sslmode= %s",host,port ,user, pas sword,dbname,sslmode)

	// connection open
	db, err := sql.Open("postgres",psqlconn)
	if err != nil{
		return nil, fmt.Printf("DB connection failed %v\n",err)
	}
	
	return db, nil
}
// insertStmnt := `insert into "members" values(7821, 'kriti', 'kriti@example.com',8873422290, '2024-07-11', 'A003');`
	// _ ,e := db.Exec(insertStmnt)
	// if e != nil{
	// 	fmt.Println(" insertion failed %v",e)
	// }