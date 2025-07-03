package snowvillage

import "github.com/go-chi/chi/v5"

func (s *snowVillageService) GetBasePath() string {
	return "/snow-village"
}

func (s *snowVillageService) GetRoutes() *chi.Mux {
	mux := chi.NewMux()

	mux.Post("/message", s.newMessage)

	return mux
}
