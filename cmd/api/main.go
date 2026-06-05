package main

import (
	"net/http"

	"github.com/common-nighthawk/go-figure"
	"github.com/go-chi/chi"
	"github.com/jetqin/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.RegisterHandlers(r)

	fig := figure.NewFigure("GO API", "", true)
	fig.Print()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error(err)
	}
}
