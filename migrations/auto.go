package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cluster := gocql.NewCluster(os.Getenv("Host"))
	cluster.Keyspace = "system"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	err = session.Query(`
        CREATE KEYSPACE IF NOT EXISTS url_shortener
        WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}
    `).Exec()
	if err != nil {
		log.Fatal(err)
	}

	session.Close()

	cluster.Keyspace = "url_shortener"
	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	err = session.Query(`
        CREATE TABLE IF NOT EXISTS schema_migrations (
            id text PRIMARY KEY
        )
    `).Exec()
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir("migrations")
	if err != nil {
		log.Fatal(err)
	}

	var migrationFiles []string
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".cql" {
			migrationFiles = append(migrationFiles, f.Name())
		}
	}
	sort.Strings(migrationFiles)

	for _, m := range migrationFiles {
		var exists string
		err := session.Query(`SELECT id FROM schema_migrations WHERE id = ?`, m).Scan(&exists)
		if err == gocql.ErrNotFound {
			content, err := os.ReadFile(filepath.Join("migrations", m))
			if err != nil {
				log.Fatalf("Reading migration error %s: %v", m, err)
			}

			err = session.Query(string(content)).Exec()
			if err != nil {
				log.Fatalf("Reading migration error %s: %v", m, err)
			}

			err = session.Query(`INSERT INTO schema_migrations (id) VALUES (?)`, m).Exec()
			if err != nil {
				log.Fatalf("Loading migration error %s: %v", m, err)
			}

			fmt.Printf("Applied migration: %s\n", m)
		} else if err != nil {
			log.Fatalf("Applying migration error %s: %v", m, err)
		} else {
			fmt.Printf("Migration already applied: %s\n", m)
		}
	}
}
