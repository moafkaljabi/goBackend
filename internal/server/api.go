package server

import (
	"context"
	"encoding/json"
	"fmt"
	"goBackend/internal/database"
	"goBackend/internal/models"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func (s *APIServer) Run() {
	router := mux.NewRouter()

	// Account routes
	router.HandleFunc("/Account", MakeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/Account/{id}", MakeHTTPHandleFunc(s.handleGetAccount))

	// Device routes
	router.HandleFunc("/Device", MakeHTTPHandleFunc(s.handleCreateDevice)).Methods(("POST"))
	server := &http.Server{
		Addr:    s.listenAddr,
		Handler: router,
	}

	// Create a channel to listen for OS signals (Ctrl+C, SIGTERM)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine so it doesn't block
	go func() {
		log.Println("JSON server running on port:", s.listenAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %s", err)
		}
	}()

	// Wait for SIGINT or SIGTERM
	<-stop
	log.Println("Shutting down server...")

	// Gracefully shut down the server with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %s", err)
	}

	log.Println("Server stopped successfully")
}

type APIServer struct {
	listenAddr string
	store      database.Storage
}

func NewAPIServer(listenAddr string, store database.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s ", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {

	id := mux.Vars(r)["id"]
	fmt.Println(id)

	return WriteJSON(w, http.StatusOK, &models.Account{})

}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransferAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

// Function signature
type ApiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

// API func to http handler func
func MakeHTTPHandleFunc(f ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
