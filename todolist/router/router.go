package router

import "net/http"

func router() {
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/createItem", createItemHandler)
	// http.HandleFunc("/getCompletedItems", getCompletedItemsHandler)
	// http.HandleFunc("/getInCompletedItems", getInCompletedItemsHandler)
	// http.HandleFunc("/updateItem", updateItemHandler)
	// http.HandleFunc("/deleteItems", deleteItemHandler)
}
