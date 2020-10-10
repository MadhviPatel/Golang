package middleware

import (
	"io"
	"net/http"
)

func topHeadLinesHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Top Headlines")
}
