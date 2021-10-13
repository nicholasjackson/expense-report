package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/env"
	"github.com/nicholasjackson/expense-report/expense/database"
	"github.com/nicholasjackson/expense-report/expense/handlers"
)

var listenAddress = env.String("LISTEN_ADDR", false, "0.0.0.0:15001", "IP address and port to bind service to")
var connectionString = env.String("MYSQL_CONNECTION", false, "root:password@tcp(127.0.0.1:3307)/DemoExpenses", "connection string for expense database")

func main() {
	options := hclog.DefaultOptions
	options.Color = hclog.AutoColor
	logger := hclog.New(options)

	err := env.Parse()
	if err != nil {
		logger.Error("Unable to parse environment variables", "error", err)
		os.Exit(1)
	}

	// create the database connection
	logger.Info("Attempting to connect to the database", "connection", *connectionString)
	db, err := database.New(*connectionString)
	if err != nil {
		logger.Error("Unable to connect to the database", "error", err)
		os.Exit(1)
	}

	ex := handlers.NewExpense(logger, db)

	r := mux.NewRouter()
	r.HandleFunc("/api/expense", ex.HandleGET).Methods(http.MethodGet)
	r.HandleFunc("/api/expense", ex.HandlePOST).Methods(http.MethodPost)
	http.Handle("/", r)

	logger.Info("Starting server on", "address", *listenAddress)

	err = http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		logger.Error("Unable to start server", "error", err)
		os.Exit(1)
	}
}
