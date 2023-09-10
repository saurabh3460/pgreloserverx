package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/lib/pq"
)

func main() {
    // Define the PostgreSQL connection string
    connStr := "host=0.0.0.0 user=username password=quote dbname=yourdb sslmode=disable"
    
    // Open a connection to the database
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Ping the database to ensure it's accessible
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    // Query for the PostgreSQL server version
    var version string
    err = db.QueryRow("SELECT version()").Scan(&version)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(version)
}

