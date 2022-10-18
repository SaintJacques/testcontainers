package main

import (
	"context"
	"log"
	"os"

	"testcntrns/db"
)

func main() {
	cfg, err := readConfig()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
		os.Exit(1)
	}

	if _, err := db.New(context.Background(), cfg.DBConnStr); err != nil {
		log.Fatalf("failed connect to db: %v", err)
		os.Exit(1)
	}
}
