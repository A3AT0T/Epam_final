package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"

	"Epam_final/pkg/app"
	"Epam_final/pkg/db"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	dbAddr := os.Getenv("DB_ADDRESS")
	if dbAddr == "" {
		log.Fatalf("empty DB_ADDRESS environment variable")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("empty PORT environment variable")
	}

	conn, err := db.GetConn(dbAddr)
	if err != nil {
		log.Fatalf("conncet to db: %s", err)
	}

	if err := db.Migrate(conn); err != nil {
		log.Fatalf("migrate db: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	api, err := app.New(conn)
	if err != nil {
		log.Fatalf("crate application api: %s", err)
	}
	errCh := make(chan error, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		if err := api.Run(ctx, port); err != nil {
			errCh <- err
		}
		wg.Done()
	}()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, os.Kill)

	select {
	case sig := <-stopCh:
		log.Printf("cought os signal [%s], shutting down application...", sig)
		cancel()
		wg.Wait()
	case err := <-errCh:
		log.Fatalf("application error: %s", err)
	}
}
