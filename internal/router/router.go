package router

import (
	"github.com/gorilla/mux"
	"github.com/knvi/modnevi/internal/middleware"
	"github.com/knvi/modnevi/internal/routes/api"
)

var Router = mux.NewRouter()

func init() {
	Router.Use(middleware.LogRequest)
	
	api.Register(Router)
}