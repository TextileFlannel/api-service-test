package handlers

import (
	"api-service-test/internal/models"
	"api-service-test/internal/service"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type AnswerHandler struct {
	service *service.Service
}

func NewAnswerHandler(s *service.Service) *AnswerHandler {
	return &AnswerHandler{service: s}
}

func (h *AnswerHandler) Create(w http.ResponseWriter, r *http.Request) {
	questionID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var req models.CreateAnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Text == "" {
		http.Error(w, "Text is required", http.StatusBadRequest)
		return
	}

	answer, err := h.service.CreateAnswer(r.Context(), uint(questionID), req.UserID, req.Text)
	if err != nil {
		if err == errors.New("not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

func (h *AnswerHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	answer, err := h.service.GetAnswer(r.Context(), uint(id))
	if err != nil {
		if err == errors.New("not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answer)
}

func (h *AnswerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	if err := h.service.DeleteAnswer(r.Context(), uint(id)); err != nil {
		if err == errors.New("not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
