package data

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/supabase-community/supabase-go"
)

type Suggestion struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Pinned      bool      `json:"pinned"`
	Status      string    `json:"status"`
	Title       string    `json:"title"`
	Upvotes     []int64   `json:"upvotes"`
	Comments    []int64   `json:"comments"`
}

func ValidateSuggestion(s *Suggestion) error {
	if strings.TrimSpace(s.Title) == "" {
		return errors.New("title is required")
	}
	if len(s.Title) < 5 || len(s.Title) > 100 {
		return errors.New("title must be between 5 and 100 characters")
	}

	if strings.TrimSpace(s.Description) == "" {
		return errors.New("description is required")
	}
	if len(s.Description) < 10 {
		return errors.New("description must be at least 10 characters")
	}

	if strings.TrimSpace(s.Category) == "" {
		return errors.New("category is required")
	}

	allowedStatuses := map[string]bool{
		"live":        true,
		"in-progress": true,
		"planned":     true,
	}
	if !allowedStatuses[strings.ToLower(s.Status)] {
		return errors.New("invalid status (must be 'live', 'in-progress', or 'planned')")
	}

	if s.Upvotes == nil {
		s.Upvotes = []int64{}
	}
	if s.Comments == nil {
		s.Comments = []int64{}
	}
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}

	return nil
}

type SuggestionModel struct {
	DB *supabase.Client
}

// Add Handler struct to bind your HTTP handler methods
type Handler struct {
	Suggestions SuggestionModel
}

// GET /suggestions
func (h *Handler) GetSuggestions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	resp, err := h.Suggestions.DB.
		From("suggestions").
		Select("*").
		Execute(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch suggestions: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var suggestions []Suggestion
	if err := json.NewDecoder(resp.Body).Decode(&suggestions); err != nil {
		http.Error(w, "Failed to decode suggestions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, suggestions)
}

// POST /suggestions
func (h *Handler) CreateSuggestion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	var s Suggestion
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := ValidateSuggestion(&s); err != nil {
		http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.Suggestions.DB.
		From("suggestions").
		Insert(s).
		Execute(ctx)
	if err != nil {
		http.Error(w, "Failed to create suggestion: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var inserted []Suggestion
	if err := json.NewDecoder(resp.Body).Decode(&inserted); err != nil {
		http.Error(w, "Failed to decode created suggestion: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(inserted) == 0 {
		http.Error(w, "No suggestion returned", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, inserted[0])
}

// DELETE /suggestions/:id
func (h *Handler) DeleteSuggestion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	idStr := ps.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid suggestion ID", http.StatusBadRequest)
		return
	}

	resp, err := h.Suggestions.DB.
		From("suggestions").
		Delete().
		Eq("id", id).
		Execute(ctx)
	if err != nil {
		http.Error(w, "Failed to delete suggestion: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		http.Error(w, "Failed to delete suggestion: "+resp.Status, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// PATCH /suggestions/:id
func (h *Handler) UpdateSuggestion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	idStr := ps.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid suggestion ID", http.StatusBadRequest)
		return
	}

	var s Suggestion
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.Suggestions.DB.
		From("suggestions").
		Update(s).
		Eq("id", id).
		Execute(ctx)
	if err != nil {
		http.Error(w, "Failed to update suggestion: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var updated []Suggestion
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		http.Error(w, "Failed to decode updated suggestion: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(updated) == 0 {
		http.Error(w, "Suggestion not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, updated[0])
}
