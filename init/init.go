package init

import (
	"circledigital.in/api/utils/common"
	"github.com/go-chi/chi/v5"
)

type app struct {
	mux *chi.Mux
}

func (a *app) GetRouter() *chi.Mux {
	return a.mux
}

func (a *app) initApplication() {
	a.mux = a.createRouter()
}

func GetApplication() common.IApp {
	appObj := &app{}
	appObj.initApplication()

	return appObj
}
