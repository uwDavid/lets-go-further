package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// config struct to hold config settings for application
// we will read in these config settings using cmd-line flags
type config struct {
	port int
	env  string
}

// app struct to hold dependencies for HTTP handler, helpers, and middleware
type application struct {
	config config
	logger *log.Logger
}

func main() {
	// initiate instance of config struct
	var cfg config

	// read in config flags
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// initialize logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// initialize app struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// server config
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start Server
	logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
