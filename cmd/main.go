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
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://football:password@localhost:5433/football?sslmode=disable", "PostgreSQL DSN")
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

	logger.Printf("Database Connection Pool Established")
	// First, call the excel handler to convert excel data into sql data.
	app.GetClubData()
	app.GetPlayerData()
	app.GetPlayerStatsData()

	// Second, perform proprocessing on the data stored in the database
	// 5 number summary for: minutes_played
	// Precomputations

	result, err := app.models.Player_Stats.GetMinutes()
	if err != nil {
		log.Fatal("Error db query")
	}
	_, q1, _, q3, _, err := app.getFiveNumberSummary(result)
	if err != nil {
		log.Fatal("Eror db query")
	}
	precomp := &Precomputations{
		MinutesLower: q1 + 0*q3,
	}

	done := make(chan interface{})
	defer close(done)

	stats, err := app.models.Player_Stats.GetAllStats()
	if err != nil {
		log.Fatal("Pipeline failed")
	}

	pipeline := app.transformed(done, app.cleaning(done, app.generator(done, stats...), *precomp), *precomp)
	for v := range pipeline {
		if v.Minutes_Played != -1 {
			app.models.Player_Stats.Update(v)
		} else {
			app.models.Player_Stats.Delete(v.Player_ID, v.Current_Club_ID, v.Season)
		}
	}

	log.Printf("Data Cleaned...")

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
