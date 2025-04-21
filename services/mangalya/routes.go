package mangalya

import "github.com/go-chi/chi/v5"

func (ms *mangalyaService) GetBasePath() string {
	return "/mangalya"
}

func (ms *mangalyaService) GetRoutes() *chi.Mux {
	mux := chi.NewMux()
	mux.Post("/form-submit", ms.newFormSubmission)
	return mux
}
