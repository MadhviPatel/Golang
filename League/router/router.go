package router

import (
	"net/http"

	"github.com/pmadhvi/Golang/League/middleware"
)

//Router contains list of all api's routes
func Router() {
	http.HandleFunc("/leagues/healthz", middleware.HealthzHandler)
	http.HandleFunc("/leagues/all", middleware.LeaguesHandler)
	http.HandleFunc("/leagues/id/{id}", middleware.LeaguesIDHandler)
}
