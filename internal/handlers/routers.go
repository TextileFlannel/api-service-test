package handlers

import (
	"api-service-test/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(s *service.Service) http.Handler {
	r := chi.NewRouter()

	questionsHandler := NewQuestionHandler(s)
	answerHandler := NewAnswerHandler(s)

	r.Route("/questions", func(r chi.Router) {
		r.Get("/", questionsHandler.List)
		r.Post("/", questionsHandler.Create)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", questionsHandler.Get)
			r.Delete("/", questionsHandler.Delete)

			r.Route("/answers", func(r chi.Router) {
				r.Post("/", answerHandler.Create)
			})
		})
	})

	r.Route("/answer", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", answerHandler.Get)
			r.Delete("/", answerHandler.Delete)
		})
	})

	return r
}
