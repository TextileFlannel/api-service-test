package handlers

import (
	"api-service-test/internal/service"
	"net/http"
)

type AnswerHandler struct {
	service *service.Service
}

func NewAnswerHandler(s *service.Service) *AnswerHandler {
	return &AnswerHandler{service: s}
}

func (h *AnswerHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *AnswerHandler) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *AnswerHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
