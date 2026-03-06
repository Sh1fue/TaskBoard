package main

import (
    "context"
    "fmt"
    "log"
	"github.com/joho/godotenv"
	"os"
    "github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	env:= godotenv.Load()
	if env != nil{
		log.Fatal(".env Пустой")
	}
	con:= fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
	)
	pool, err := pgxpool.New(context.Background(), con)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
    defer pool.Close()

    var greeting string
    err = pool.QueryRow(context.Background(), "SELECT 'Hello, pgx with .env!'").Scan(&greeting)
    if err != nil {
        log.Fatalf("QueryRow failed: %v\n", err)
    }

    fmt.Println(greeting)
}
