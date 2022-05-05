package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "0.1.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn           string
		maxOpensConns int
		maxIdleConns  int
		maxIdleTime   string
	}
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// Read command-line flags
	flag.IntVar(&cfg.port, "port", 4000, "API listen port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stag|prod)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "PostgreSQL DSN")

	// Read connection pool options
	flag.IntVar(&cfg.db.maxOpensConns, "db-max-open-conns", 25, "PSQL max opens connection")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PSQL max idele connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PSQL max connection idle time")

	flag.Parse()

	// Create stdout logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Connect to DB
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	logger.Println("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
	}

	// set server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// Set maximum number of open (in-use + idle) connections in pool
	//passing 0 value mean there is no limit
	db.SetMaxOpenConns(cfg.db.maxOpensConns)

	// Set maximum number of idle connection in pool (0 - no limit)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	// Use time.ParseDuration() to convert string to idle timeout
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	//Create context with 5-second deadline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
