package main

import (
	"bsa-core/pkg/api"
	"bsa-core/pkg/state"
	"log"
	"net/http"
	"os"
)

func main() {
	// 1. Initialize Logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting BSA Core Service...")

	// 2. Initialize BSA State Manager
	repoPath := os.Getenv("BSA_REPO_PATH")
	if repoPath == "" {
		repoPath = "./data" // Default to local data directory
	}
	bsa := state.NewBSA(repoPath)

	// 3. Start Reconciliation Loop (Async)
	go bsa.Reconcile()
	log.Println("Reconciliation loop started.")

	// 4. Initialize API Handler
	handler := api.NewHandler(bsa)

	// 5. Setup Router
	http.HandleFunc("/api/v1/state", handler.GetState)
	http.HandleFunc("/api/v1/propose", handler.ProposeChange)

	// 6. Start HTTP Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
