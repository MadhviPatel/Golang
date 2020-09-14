package router

import (
	"net/http"
	"github.com/pmadhvi/Golang/todolist/middleware"
) 

func Router() {
	http.HandleFunc("/healthz", middleware.healthzHandler)
	http.HandleFunc("/createItem", middleware.createItemHandler)
	// http.HandleFunc("/getCompletedItems", middleware.getCompletedItemsHandler)
	// http.HandleFunc("/getInCompletedItems", middleware.getInCompletedItemsHandler)
	// http.HandleFunc("/updateItem", middleware.updateItemHandler)
	// http.HandleFunc("/deleteItems", middleware.deleteItemHandler)
}
