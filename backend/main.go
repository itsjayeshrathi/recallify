package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/itsjayeshrathi/recallify/internal/app"
	"github.com/itsjayeshrathi/recallify/internal/routes"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error laoding .env file")
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT is not set in the .env file")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Error converting PORT to integer: %v\n", err)
	}

	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	r := routes.SetupRouts(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
