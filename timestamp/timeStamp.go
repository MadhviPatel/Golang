package main

import (
	"fmt"
	"net/http"
	"time"
)

func currentTimestamp(rw http.ResponseWriter, req *http.Request) {
	t := time.Now()
	rw.Write([]byte(t.String()))
}

func main() {
	fmt.Println("Inside Main")
	http.HandleFunc("/api/timestamp", currentTimestamp)
	//http.HandleFunc("/api/timestamp/date/{date}", timestamp)

	http.ListenAndServe(":9000", nil)
}
