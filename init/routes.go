package init

import (
	"circledigital.in/api/services/cozy"
	"circledigital.in/api/utils/common"
	"circledigital.in/api/utils/custom"
	"circledigital.in/api/utils/payload"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"net/http"
	"time"
)

// serviceFactory defines type for getting a service
type serviceFactory func(app common.IApp) common.IService

var services = []serviceFactory{
	cozy.CreateCozyService,
}

// handle400 returns custom responses for not found routes and not allowed methods
func (a *app) handle400(router *chi.Mux) {
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		err := &custom.RequestError{
			Status:  http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
		payload.HandleError(w, err)
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		err := &custom.RequestError{
			Status:  http.StatusMethodNotAllowed,
			Message: http.StatusText(http.StatusMethodNotAllowed),
		}
		payload.HandleError(w, err)
	})
}

func (a *app) createRouter() *chi.Mux {
	mux := chi.NewMux()

	// application middlewares
	mux.Use(httprate.LimitByIP(100, time.Minute))
	mux.Use(middleware.Heartbeat("/"))
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.StripSlashes)
	mux.Use(middleware.AllowContentType("application/json"))

	a.handle400(mux)

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// add services routes
	for _, factory := range services {
		service := factory(a)
		mux.Mount(service.GetBasePath(), service.GetRoutes())
	}

	return mux
}
