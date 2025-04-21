package cozy

import "github.com/go-chi/chi/v5"

func (cs *cozyService) GetBasePath() string {
	return "/cozy"
}

func (cs *cozyService) GetRoutes() *chi.Mux {
	mux := chi.NewMux()

	mux.Post("/reservation", cs.addNewReservation)

	return mux
}
