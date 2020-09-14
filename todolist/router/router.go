package router

import (
	"net/http"

	"github.com/pmadhvi/Golang/todolist/middleware"
)

func Router() {
	http.HandleFunc("/healthz", middleware.HealthzHandler)
	http.HandleFunc("/createItem", middleware.CreateItemHandler)
	http.HandleFunc("/getCompletedItems", middleware.GetCompletedItemsHandler)
	http.HandleFunc("/getInCompletedItems", middleware.GetInCompletedItemsHandler)
	http.HandleFunc("/updateItem", middleware.UpdateItemHandler)
	http.HandleFunc("/deleteItems", middleware.DeleteItemHandler)
}
