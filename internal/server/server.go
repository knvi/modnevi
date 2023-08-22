package server

import (
	"net/http"
	"time"

	"github.com/knvi/modnevi/internal/router"
)

func NewServer(addr string) error {
	srv := &http.Server{
		Handler: router.Router,
		Addr:    addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	return srv.ListenAndServe()
}