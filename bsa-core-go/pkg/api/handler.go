package api

import (
	"bsa-core/pkg/state"
	"encoding/json"
	"net/http"
)

// Handler holds dependencies for API handlers.
type Handler struct {
	bsa *state.BSA
}

// NewHandler creates a new Handler instance.
func NewHandler(bsa *state.BSA) *Handler {
	return &Handler{bsa: bsa}
}

// GetState handles GET /api/v1/state
func (h *Handler) GetState(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	version := r.URL.Query().Get("version")
	// If version is empty, it defaults to latest in the underlying implementation

	s, err := h.bsa.GetState(version)
	if err != nil {
		http.Error(w, "Failed to get state", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

// ProposeChange handles POST /api/v1/propose
func (h *Handler) ProposeChange(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var proposal state.Proposal
	if err := json.NewDecoder(r.Body).Decode(&proposal); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.bsa.ProposeChange(proposal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"status": "accepted", "intent_id": proposal.IntentID})
}
