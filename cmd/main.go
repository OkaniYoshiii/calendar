package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/OkaniYoshiii/calendar/internal/database"
	"github.com/OkaniYoshiii/calendar/internal/handlers"
	"github.com/OkaniYoshiii/calendar/internal/repository"
	"github.com/joho/godotenv"
)

var address = flag.String("address", ":8080", "TCP address the server will listen to")

func main() {
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("\"DATABASE_DSN\" environment variable has not been set")
	}

	db, err := database.Connect(dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	queries := repository.New(db)

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./website/dist"))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", fs))
	mux.Handle("GET /{$}", &handlers.HomeHandler{
		Queries: queries,
	})

	server := http.Server{
		Addr:              *address,
		Handler:           mux,
		ReadTimeout:       time.Second * 15,
		ReadHeaderTimeout: time.Second * 1,
		WriteTimeout:      time.Second,
		IdleTimeout:       time.Second,
		MaxHeaderBytes:    1000,
	}

	fmt.Printf("Server listening on address : \"%s\"\n", *address)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
