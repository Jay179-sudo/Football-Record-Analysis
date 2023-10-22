package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"time"

	"github.com/Jay179-sudo/FootballRecordAnalysis/internal/data"
	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	db struct {
		dsn string
	}
}

type application struct {
	logger *log.Logger
	models data.Models
}

func main() {
	var cfg config
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://admin:password@localhost:5433/football?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	app := &application{
		logger: logger,
		models: data.NewModels(db),
	}

	app.printState(logger)
	logger.Printf("Database Connection Pool Established")
	// First, call the excel handler to convert excel data into sql data.
	// Second, perform proprocessing on the data stored in the database
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (a *application) printState(logger *log.Logger) {
	logger.Printf("Hi!")
}
