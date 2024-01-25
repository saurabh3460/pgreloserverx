package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var (
	sourceDB      string
	destinationDB string
	dbDriver      = "postgres"
)

func init() {
	flag.StringVar(&sourceDB, "source", "", "source database connection string")
	flag.StringVar(&destinationDB, "dest", "", "destination database connection string")
}

func main() {

	flag.Parse()

	// Validate that sourceDB and destinationDB are not empty
	if sourceDB == "" || destinationDB == "" {
		log.Fatal("Both source and destination database connection strings must be provided.")
	}

	// Connect to the source database
	sourceDBConn, err := sql.Open(dbDriver, sourceDB)
	if err != nil {
		log.Fatal("Error connecting to the source database:", err)
	}
	defer sourceDBConn.Close()

	// Connect to the destination database
	destinationDBConn, err := sql.Open(dbDriver, destinationDB)
	if err != nil {
		log.Fatal("Error connecting to the destination database:", err)
	}
	defer destinationDBConn.Close()

	var timeout time.Duration = 5

	sctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	dbConnectionCheck(sctx, sourceDBConn)

	dctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	dbConnectionCheck(dctx, destinationDBConn)

	// Query and log the version for the source database
	sourceVersion, err := queryDataaseVersion(sourceDBConn)
	if err != nil {
		log.Fatal("Error querying source database version:", err)
	}
	log.Printf("Source Database Version: %s", sourceVersion)

	// Query and log the Size of the source database
	sourceSize, err := queryDatabaseSize(sourceDBConn)
	if err != nil {
		log.Fatal("Error querying source database size:", err)
	}
	log.Printf("Source Database Size in Bytes: %d", sourceSize)

	// Query and log the version for the destination database
	destinationVersion, err := queryDataaseVersion(destinationDBConn)
	if err != nil {
		log.Fatal("Error querying destination database version:", err)
	}
	log.Printf("Destination Database Version: %s", destinationVersion)

}

func dbConnectionCheck(ctx context.Context, db *sql.DB) {
	ticker := time.NewTicker(1000 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Fatal("Context timeout")
		case <-ticker.C:
			err := db.Ping()
			if err == nil {
				log.Println("connected")
				return
			}
			log.Printf("Try connecting... Error: %v", err)
		}
	}
}

func queryDataaseVersion(db *sql.DB) (string, error) {
	var version string
	err := db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		return "", err
	}
	return version, nil
}

func queryDatabaseSize(db *sql.DB) (int64, error) {
	var size int64
	err := db.QueryRow("SELECT pg_database_size(current_database());").Scan(&size)
	if err != nil {
		return 0, err
	}
	return size, nil
}
