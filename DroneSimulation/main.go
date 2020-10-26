package main

import (
	"fmt"
	"net/http"

	"github.com/pmadhvi/Golang/DroneSimulation/router"
)

func main() {
	fmt.Println("This is a Drone Simulator Server")

	router.Router()
	http.ListenAndServe(":4000", nil)
}
