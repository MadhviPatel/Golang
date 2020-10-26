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
