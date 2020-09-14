package middleware

import (
	"io"
	"net/http"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// const (
// 	DB_USER = "madhvip"
// 	DB_PASS = "madhvip"
// 	DB_NAME = "friends_repo"
// )

// var db *sql.DB
// var err error

// func init() {
// 	log.Info("Connecting to todolist database")
// 	db, err = connectToDatabase()
// 	chkerr(err)
// }

// type response struct {
// 	ID      int    `json:"id,omitempty"`
// 	Message string `json:"message,omitempty"`
// }

// func connectToDatabase() (db *sql.DB, err error) {
// 	dbinfo := fmt.Sprintf("host=localhost port= 5432 user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
// 	db, err = sql.Open("postgres", dbinfo)
// 	return
// }

// func chkerr(err error) {
// 	if err != nil {
// 		log.Fatal(err.Error())
// 		panic(err)
// 	}
// }

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

// func createItemHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Info("createItem API")
// 	w.Header().Set("Content-Type", "application/json")

// 	var todoitem models.TodoItemModel
// 	err := json.NewDecoder(r.Body).Decode(&todoitem)
// 	if err != nil {
// 		log.Fatalf("Unable to decode the request body.  %v", err)
// 	}

// 	insertedId := insert(todoitem)
// 	resp := response{ID: insertedId, Message: "Item created successfully"}
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(resp)
// }

// func insert(todoitem models.TodoItemModel) (lastId int) {
// 	log.Info("Inserting to db")
// 	err := db.QueryRow("insert into todolist(description, created, completed) values ($1,$2,$3) returning id;", todoitem.description, todoitem.created, todoitem.completed).Scan(&lastId)
// 	chkerr(err)
// 	log.Infof("createdItemId: %d", lastId)
// 	return
// }
