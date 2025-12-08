package state

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Proposal represents a change request from an agent.
type Proposal struct {
	IntentID    string `json:"intent_id"`
	DeltaPatch  string `json:"delta_patch"`
	Provisional bool   `json:"provisional"`
}

// BSA (Deterministic State Authority) manages the canonical state and enforces integrity.
type BSA struct {
	RepoPath  string
	liveIndex map[string]interface{}
	mu        sync.RWMutex
}

// NewBSA initializes a new BSA instance.
func NewBSA(repoPath string) *BSA {
	return &BSA{
		RepoPath:  repoPath,
		liveIndex: make(map[string]interface{}),
	}
}

// GetState retrieves the canonical state for a given version.
// For now, it returns the current live index. In a real implementation,
// it would checkout the specific git version.
func (b *BSA) GetState(version string) (map[string]interface{}, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	// Deep copy to avoid race conditions if the caller modifies the map
	stateCopy := make(map[string]interface{})
	for k, v := range b.liveIndex {
		stateCopy[k] = v
	}

	return stateCopy, nil
}

// ProposeChange validates a proposal and writes it to a staging area.
func (b *BSA) ProposeChange(proposal Proposal) error {
	if proposal.IntentID == "" {
		return fmt.Errorf("intent_id is required")
	}
	if proposal.DeltaPatch == "" {
		return fmt.Errorf("delta_patch is required")
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	// In a real implementation, this would apply the patch to a staging branch.
	// For this core implementation, we simulate staging by logging and updating the in-memory index temporarily
	// if it's not provisional, or just acknowledging it.

	log.Printf("Received proposal: %s (Provisional: %v)", proposal.IntentID, proposal.Provisional)

	// Simulate writing to staging area (e.g., a file or a git branch)
	stagingFile := filepath.Join(b.RepoPath, "staging", fmt.Sprintf("%s.json", proposal.IntentID))
	if err := os.MkdirAll(filepath.Dir(stagingFile), 0755); err != nil {
		return fmt.Errorf("failed to create staging directory: %w", err)
	}

	data, err := json.Marshal(proposal)
	if err != nil {
		return fmt.Errorf("failed to marshal proposal: %w", err)
	}

	if err := os.WriteFile(stagingFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write proposal to staging: %w", err)
	}

	return nil
}

// Reconcile runs the hard sync logic.
// It pulls the latest canonical state, validates proposals, and commits changes.
func (b *BSA) Reconcile() {
	ticker := time.NewTicker(5 * time.Second) // Run every 5 seconds
	defer ticker.Stop()

	for range ticker.C {
		b.runReconciliationLoop()
	}
}

func (b *BSA) runReconciliationLoop() {
	b.mu.Lock()
	defer b.mu.Unlock()

	// 1. Pull latest canonical state from main branch (Simulated)
	// log.Println("Reconciling: Pulling latest state...")

	// 2. Read proposals from staging
	stagingDir := filepath.Join(b.RepoPath, "staging")
	files, err := os.ReadDir(stagingDir)
	if err != nil {
		// Staging might not exist yet
		return
	}

	if len(files) == 0 {
		return
	}

	log.Printf("Reconciling %d proposals...", len(files))

	// 3. Validate and Apply (Simulated atomic commit)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Read proposal
		content, err := os.ReadFile(filepath.Join(stagingDir, file.Name()))
		if err != nil {
			log.Printf("Failed to read proposal %s: %v", file.Name(), err)
			continue
		}

		var prop Proposal
		if err := json.Unmarshal(content, &prop); err != nil {
			log.Printf("Failed to unmarshal proposal %s: %v", file.Name(), err)
			continue
		}

		// Apply to live index (Simulating state update)
		// In reality, this would apply the delta patch.
		// Here we just store the proposal ID as a "change" in the state for demonstration.
		b.liveIndex["last_processed_intent"] = prop.IntentID
		b.liveIndex[prop.IntentID] = "applied"

		// Remove from staging
		if err := os.Remove(filepath.Join(stagingDir, file.Name())); err != nil {
			log.Printf("Failed to remove proposal %s from staging: %v", file.Name(), err)
		}
	}

	// 4. Commit to main branch (Simulated)
	// log.Println("Reconciliation complete: State updated.")
}
