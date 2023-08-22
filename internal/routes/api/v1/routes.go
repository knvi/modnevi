package v1

import (
	"github.com/gorilla/mux"
	"github.com/knvi/modnevi/internal/routes/api/v1/mod"
)

func Register(router *mux.Router) {
	subRouter := router.PathPrefix("/v1").Subrouter()

	subRouter.HandleFunc("/mod", mod.Handler)
}