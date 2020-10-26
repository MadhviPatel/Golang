package middleware

import (
	"fmt"
	"io"
	"net/http"
	log "github.com/sirupsen/logrus"
)


// HealthzHandler is a health api
func HealthzHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside Healthz handler")
	log.Info("API Health is OK")
	res.Header().Set("Content-Type", "application/json")
	io.WriteString(res, `{"alive": true}`)
}

//LeaguesHandler is an api to get all the leagues
func LeaguesHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside Leagues handler")
	res.Header().Set("Content-Type", "application/json")
	io.WriteString(res, `{"leagues": all}`)
}

//LeaguesIDHandler is an api to get leagues for an id
func LeaguesIDHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside Leagues Id handler")
	res.Header().Set("Content-Type", "application/json")
	io.WriteString(res, `{"leagues": id}`)
}




















