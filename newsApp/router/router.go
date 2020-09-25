package router

import (
	"net/http"
	m "github.com/pmadhvi/Golang/newsApp/middleware"
)

func Routes() {

	http.HandleFunc("/topHeadLines", m.topHeadLinesHandler)
}
