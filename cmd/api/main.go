package main

import (
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	apphttp "golang-test-task/internal/app/http"
	"golang-test-task/internal/app/storage"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is required")
	}

	db, err := storage.OpenPostgres(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := storage.Migrate(db); err != nil {
		log.Fatal(err)
	}

	repo := storage.NewPostgresRepo(db)
	h := apphttp.NewHandler(repo)
	router := apphttp.Router(h)

	addr := ":8080"
	srv := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("listening on %s", addr)
	log.Fatal(srv.ListenAndServe())
}
