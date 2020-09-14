package main

import (
	"net/http"

	"github.com/pmadhvi/Golang/todolist/router"
	log "github.com/sirupsen/logrus"
)

func main() {
	//defer db.Close()
	log.Info("Starting todolist server")
	router.Router()
	http.ListenAndServe(":9000", nil)

}
