package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var dbConn = createDB("postgres", "postgres://root:password@localhost:5432/doodle-guessr?sslmode=disable")

func createDB(driverName string, dataSourceName string) *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println("Error (CreateDB) when opening database connection")
		log.Printf("   |_ %v\n", err.Error())
		log.Fatal("TERMINATED")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		log.Println("Error (CreateDB) when pinging the database")
		log.Printf("   |_ %v\n", err.Error())
		log.Fatal("TERMINATED")
		return nil
	}

	return db
}

func DB() *sql.DB {
	return dbConn
}
