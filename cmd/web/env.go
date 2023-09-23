package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	csrf struct {
		key    string
		secure bool
	}
	gmaps struct {
		apiKey string
	}
}

func loadEnv() (config, error) {
	var cfg config

	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}

	cfg.port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return cfg, err
	}
	cfg.env = os.Getenv("ENV")

	cfg.db.dsn = os.Getenv("PSQL_DSN")
	cfg.db.maxIdleTime = os.Getenv("MAX_IDLE_TIME")
	cfg.db.maxOpenConns, err = strconv.Atoi(os.Getenv("MAX_OPEN_CONNS"))
	if err != nil {
		return cfg, err
	}
	cfg.db.maxIdleConns, err = strconv.Atoi(os.Getenv("MAX_IDLE_CONNS"))
	if err != nil {
		return cfg, err
	}

	cfg.csrf.key = os.Getenv("CSRF_KEY")
	cfg.csrf.secure = os.Getenv("CSRF_SECURE") == "true"

	cfg.gmaps.apiKey = os.Getenv("GMAP_API_KEY")

	return cfg, nil
}
