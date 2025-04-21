package common

import "github.com/go-chi/chi/v5"

// package common handles all the methods required by multiple services

// IService is implemented by all services and return routes exposed by the service
type IService interface {
	GetRoutes() *chi.Mux
	GetBasePath() string
}

// IApp is an application interface with all the configurations
type IApp interface {
	GetRouter() *chi.Mux
}
