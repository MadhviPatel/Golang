package main

import (
	"fmt"
	"net/http"

	"github.com/pmadhvi/Golang/newsApp/router"
)

func main() {
	fmt.Println("Starting NewsApp server")
	router.Routes()
	http.ListenAndServe(":9000", nil)

}
