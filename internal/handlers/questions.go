package handlers

import (
	"api-service-test/internal/service"
	"net/http"
)

type QuestionHandler struct {
	service *service.Service
}

func NewQuestionHandler(s *service.Service) *QuestionHandler {
	return &QuestionHandler{service: s}
}

func (h *QuestionHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *QuestionHandler) List(w http.ResponseWriter, r *http.Request) {

}

func (h *QuestionHandler) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *QuestionHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
