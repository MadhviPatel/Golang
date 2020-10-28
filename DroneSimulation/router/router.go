package router

import (
	"net/http"

	"github.com/pmadhvi/Golang/DroneSimulation/middleware"
)

//Router contains list of all api's routes
func Router() {
	http.HandleFunc("/drones/healthz", middleware.HealthzHandler)
	// http.HandleFunc("/drones/all", middleware.LeaguesHandler)
	// http.HandleFunc("/drones/id/{id}", middleware.LeaguesIDHandler)
}
