package db

import (
    "fmt"
    "log"
    "os"
    "database/sql"
    "github.com/joho/godotenv"   // package used to read the .env file
    _ "github.com/lib/pq"      // postgres golang driver
)

const (
    PORT = 5432
)


// create connection with postgres db
func createDBConnection() *sql.DB {
    // load .env file
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    port := PORT
    // access variables from .env file
    host, username, password, database :=
    os.Getenv("POSTGRES_HOST"),
    os.Getenv("POSTGRES_USER"),
    os.Getenv("POSTGRES_PASSWORD"),
    os.Getenv("POSTGRES_DB")

    db_conn_str := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, username, password, database)
    // Open the connection
    db, err := sql.Open("postgres", db_conn_str)
    if err != nil {
        panic(err)
    }
    // check the connection
    err = db.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("Successfully connected!")
    // return the connection
    return db
}