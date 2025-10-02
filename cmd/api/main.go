package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bujdoluk/feedback-api/internal/data"
	"github.com/bujdoluk/feedback-api/internal/jsonlog"
	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

var (
	version string
)

type config struct {
	port        int
	env         string
	supabaseURL string
	supabaseKey string
}

type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
}

func main() {
	var cfg config

	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found (using system env vars)")
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_PUBLIC_KEY")

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	if supabaseURL == "" || supabaseKey == "" {
		logger.PrintFatal(fmt.Errorf("SUPABASE_URL or SUPABASE_PUBLIC_KEY not set"), nil)
	}

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Enviroment (development|staging|production)")
	flag.StringVar(&cfg.supabaseURL, "url", supabaseURL, "Supabase URL")
	flag.StringVar(&cfg.supabaseKey, "key", supabaseKey, "Supabase public key")

	flag.Parse()

	client, err := supabase.NewClient(cfg.supabaseURL, cfg.supabaseKey, nil)
	if err != nil {
		logger.PrintFatal(fmt.Errorf("failed to initialize Supabase client: %w", err), nil)
	}

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(client),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
