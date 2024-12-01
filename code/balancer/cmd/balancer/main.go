package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"ttb-service/internal/balancer"
	"ttb-service/internal/events"
	"ttb-service/internal/handlers"
	"ttb-service/internal/utils"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	// Load the config
	err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Set up logger
	logger := utils.NewLogger()

	// Create HTTP client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Initialize HTTP Publisher
	eventPublisher := events.NewHTTPPublisher(
		viper.GetString("eventPublisher.url"), // URL from config
		client,
		logger,
	)

	// Initialize the AlgorithmService
	algorithmService := balancer.NewAlgorithmService(logger)

	// Create the BalancerHandler
	balancerHandler := handlers.NewAlgorithmBalancerHandler(algorithmService, logger, eventPublisher)
	eventHandler := handlers.NewEventHandler(eventPublisher, logger)
	statusHandler := handlers.NewHealthCheckHandler()

	// Initialize HTTP router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/balancer/apply", balancerHandler.ApplyBalancerAlgorithm).Methods("POST")
	r.HandleFunc("/balancer/events/", eventHandler.HandleEvent).Methods("POST")
	r.HandleFunc("/status", statusHandler.CheckHealth)

	// Start the HTTP server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", viper.GetString("server.port")),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Channel for handling OS interrupts
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on port %s...\n", viper.GetString("server.port"))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait for an interrupt signal (SIGINT or SIGTERM)
	<-stop
	log.Println("Received interrupt signal, shutting down gracefully...")

	// Create a deadline to wait for pending requests to finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped gracefully")
}

// loadConfig loads configuration from config.yaml
func loadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	return nil
}
