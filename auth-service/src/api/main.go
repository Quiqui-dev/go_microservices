package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Quiqui-dev/auth-service/internal/database"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

const webPort = "80"

var counts int64

type Config struct {
	DB *database.Queries
}

func main() {

	log.Println("Starting auth service")

	// todo: connect to db
	conn := connectToDB()

	if conn == nil {
		log.Panic("Can't connect to postgres")
	}

	// todo: set up config
	app := Config{
		DB: conn,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}

}

func openDB(dsn string) (*database.Queries, error) {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return database.New(db), nil
}

func connectToDB() *database.Queries {

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)

		if err != nil {
			log.Println(err)
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connected to Postgres")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for 2 seconds ....")
		time.Sleep(2 * time.Second)
		continue
	}
}
