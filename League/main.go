package main

import (
	"fmt"
	"net/http"

	"github.com/pmadhvi/Golang/League/router"
)

func main() {
	fmt.Println("This is a League Server")

	router.Router()
	http.ListenAndServe(":3000", nil)
}
