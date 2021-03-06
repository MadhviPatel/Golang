package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/pmadhvi/Golang/todolist/models"
	log "github.com/sirupsen/logrus"
)

const (
	DB_USER = "madhvip"
	DB_PASS = "madhvip"
	DB_NAME = "friends_repo"
)

var db *sql.DB
var err error

func init() {
	log.Info("Trying to connect to todolist database")
	db, err = connectToDatabase()
	chkerr(err)
}

type response struct {
	ID      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func connectToDatabase() (db *sql.DB, err error) {
	dbinfo := fmt.Sprintf("host=localhost port= 5432 user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err = sql.Open("postgres", dbinfo)
	chkerr(err)
	log.Info("Connect to todolist database")
	return
}

func chkerr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("createItem API")
	w.Header().Set("Content-Type", "application/json")

	var todoitem models.TodoItemModel
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Unable to read the request body.  %v", err)
	}
	defer r.Body.Close()
	log.Infof("req body is %s::\n", string(body))
	err = json.Unmarshal(body, &todoitem)
	if err != nil {
		log.Fatalf("Unable to unmarshall the request body.  %v", err)
	}
	insertedID := insert(todoitem)
	resp := response{ID: insertedID, Message: "Item created successfully"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}

func GetCompletedItemsHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("get completed items API")
	w.Header().Set("Content-Type", "application/json")
	completed := r.URL.Query().Get("completed")
	todoitems, err := get(completed)
	chkerr(err)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todoitems)
}

func GetInCompletedItemsHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("get incompleted items API")
	w.Header().Set("Content-Type", "application/json")
	incompleted := r.URL.Query().Get("in-completed")
	todoitems, err := get(incompleted)
	chkerr(err)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todoitems)
}

func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("update item API")
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	chkerr(err)
	var todoitem *models.TodoItemModel
	todoitem, err = getItemById(id)
	if err != nil {
		log.Printf("Item with id %d not found\n", id)
	}
	todoitem.Description = "updated_desc"
	log.Info(todoitem)
	updated, err := update(todoitem)
	chkerr(err)
	fmt.Printf("Successfully updated %d items", updated)

	resp := response{ID: todoitem.Id, Message: "Item updated successfully"}
	json.NewEncoder(w).Encode(resp)
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("delete item API")
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Errorf("Error while converting string to integer")
	}
	deleted, err := delete(id)
	chkerr(err)
	fmt.Printf("Successfully deleted %d items", deleted)
	w.WriteHeader(http.StatusOK)
	resp := response{ID: id, Message: "Item deleted successfully"}
	json.NewEncoder(w).Encode(resp)
}

func insert(todoitem models.TodoItemModel) (lastId int) {
	log.Info("Inserting data to database")
	err := db.QueryRow("insert into todolist(description, created, completed) values ($1,$2,$3) returning id;", todoitem.Description, todoitem.Created, todoitem.Completed).Scan(&lastId)
	chkerr(err)
	log.Infof("createdItemId: %d", lastId)
	return
}

func delete(id int) (affected int64, err error) {
	log.Info("deleting item from db")
	stmt, err := db.Prepare("delete from todolist where id=$1")
	chkerr(err)
	res, err := stmt.Exec(id)
	chkerr(err)
	affected, err = res.RowsAffected()
	chkerr(err)
	return
}

func update(todoitems *models.TodoItemModel) (updated int64, err error) {
	log.Info("updating item into db")
	stmt, err := db.Prepare("update todolist set description=$1 where id=$2")
	chkerr(err)
	res, err := stmt.Exec(todoitems.Description, todoitems.Id)
	chkerr(err)
	updated, err = res.RowsAffected()
	chkerr(err)
	return
}

func getItemById(id int) (itemModel *models.TodoItemModel, err error) {
	log.Info("get item by Id from db")
	row := db.QueryRow("select * from todolist where id=$1", id)
	var item models.TodoItemModel
	err = row.Scan(&item.Id, &item.Description, &item.Created, &item.Completed)
	chkerr(err)
	itemModel = &item
	return
}

func get(complete string) (todoitems []models.TodoItemModel, err error) {
	log.Info("query returning items from db")
	rows, err := db.Query("select * from todolist where completed=$1", complete)
	defer rows.Close()
	chkerr(err)

	for rows.Next() {
		var todoitem models.TodoItemModel
		err = rows.Scan(&todoitem.Id, &todoitem.Description, &todoitem.Created, &todoitem.Completed)
		chkerr(err)
		todoitems = append(todoitems, todoitem)
	}
	return
}
