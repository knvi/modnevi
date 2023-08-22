package router

import (
	"github.com/gorilla/mux"
	"github.com/knvi/modnevi/internal/routes/api"
)

var Router = mux.NewRouter()

func init() {
	api.Register(Router)
}