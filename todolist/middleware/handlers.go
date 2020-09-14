func healthzHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func createItemHandler(w http.ResponseWriter, r *http.Request) {

	var lastId int
	log.Info("createItem API")
	w.Header().Set("Content-Type", "application/json")

	log.Info("insert into db way 1")
	err := db.QueryRow("insert into todolist(description, created, completed) values ($1,$2,$3) returning id;", "description4", "2020-09-14", true).Scan(&lastId)
	chkerr(err)
	w.WriteHeader(http.StatusCreated)
	log.Infof("createdItemId: %d", lastId)
	json.NewEncoder(w).Encode(lastId)

	log.Info("insert into db way 2 using models")
	todo := &TodoItemModel{Description: "item description", Created: time.Now(), Completed: false}
	err := db.QueryRow("insert into todolist(description, created, completed) values ($1,$2,$3) returning id;", "description4", "2020-09-14", true).Scan(&lastId)
	chkerr(err)
	w.WriteHeader(http.StatusCreated)
	log.Infof("createdItemId: %d", lastId)
	json.NewEncoder(w).Encode(lastId)
}