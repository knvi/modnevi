package api

import (
	"github.com/gorilla/mux"
	v1 "github.com/knvi/modnevi/internal/routes/api/v1"
)

func Register(router *mux.Router) {
	prefix := router.PathPrefix("/api").Subrouter()

	v1.Register(prefix)
}