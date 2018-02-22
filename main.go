package main

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"database/sql"
	"github.com/mattn/go-sqlite3"
	"io/ioutil"
	"strconv"
)

// Ignore unused variable trick
var _ = sqlite3.SQLITE_DELETE

// Global
var db *sql.DB

func main() {
	// Initialize DB
	fmt.Print("Initializing database..")
	var err error
	db, err = sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared")
	handleException(err)
	query, err := db.Exec("CREATE TABLE `sneakers`( `id` INTEGER PRIMARY KEY AUTOINCREMENT, `name` VARCHAR(64) NOT NULL); CREATE TABLE `true_to_size` ( `id` INTEGER PRIMARY KEY AUTOINCREMENT, `sneaker_id` INTEGER NOT NULL, `size` INTEGER NOT NULL, FOREIGN KEY (sneaker_id) REFERENCES sneakers(id) ); INSERT INTO `sneakers` (name) VALUES (\"Jordan 3 Retro Black Cement (2018)\"), (\"Jordan 1 Retro High NRG Patent Gold Toe\"), (\"Air Foamposite One Big Bang\"), (\"adidas Yeezy Boost 350 V2 Blue Tint\"); INSERT INTO `true_to_size` (sneaker_id, size) VALUES (1, 1), (1, 2), (1, 2), (1, 3), (1, 2), (1, 3), (1, 2), (1, 2), (1, 3), (1, 4), (1, 2), (1, 5), (1, 2), (1, 3), (2, 1), (2, 2), (2, 2), (2, 3), (2, 2), (2, 3), (2, 2), (2, 2), (2, 3), (2, 4), (2, 2), (2, 5), (2, 2), (2, 3), (2, 2), (3, 3), (3, 1), (3, 3), (3, 1), (3, 2), (4, 5), (4, 5), (4, 5), (4, 5), (4, 5);")
	handleException(err)
	var _ = query
	fmt.Println("Done.")

	// Create router
	router := mux.NewRouter()

	// Register endpoints
	router.Use(JsonMiddleware)
	router.HandleFunc("/sneakers", GetSneakers).Methods("GET")
	router.HandleFunc("/sneaker/{id}", GetSneaker).Methods("GET")
	router.HandleFunc("/sneaker", CreateSneaker).Methods("POST")
	router.HandleFunc("/sneaker/{id}", DeleteSneaker).Methods("DELETE")
	router.HandleFunc("/sneaker/{id}/true-to-size", CreateTrueToSize).Methods("POST")

	// Serve HTTP
	log.Fatal(http.ListenAndServe(":8001", router))
}

func JsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

type Sneaker struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	TrueToSize string `json:"truetosize,omitempty"`
}

type TrueToSize struct {
	ID string `json:"id,omitempty"`
	SneakerID string `json:"sneakerid,omitempty"`
	Size string `json:"size,omitempty"`
}

func GetSneakers(w http.ResponseWriter, r*http.Request) {
	rows, err := db.Query("SELECT id, name, (SELECT ROUND(AVG(size), 13) FROM true_to_size WHERE true_to_size.sneaker_id = sneakers.id) as truetosize from sneakers")
	handleException(err)

	var sneakers []Sneaker

	for rows.Next() {
		var sneaker Sneaker
		rows.Scan(&sneaker.ID, &sneaker.Name, &sneaker.TrueToSize)
		sneakers = append(sneakers, sneaker)
	}

	rows.Close()

	json.NewEncoder(w).Encode(sneakers)
}

func GetSneaker(w http.ResponseWriter, r*http.Request) {
	query := fmt.Sprintf("%s%d%s", "SELECT id, name, (SELECT ROUND(AVG(size), 13) FROM true_to_size WHERE true_to_size.sneaker_id = sneakers.id) as truetosize from sneakers WHERE sneakers.id = ", 1, " LIMIT 1")
	rows, err := db.Query(query)
	handleException(err)

	var sneaker Sneaker

	for rows.Next() {
		rows.Scan(&sneaker.ID, &sneaker.Name, &sneaker.TrueToSize)
	}

	rows.Close()

	json.NewEncoder(w).Encode(sneaker)
}

func CreateSneaker(w http.ResponseWriter, r*http.Request) {
	var sneaker Sneaker
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &sneaker)

	query, err := db.Prepare("INSERT INTO sneakers (name) VALUES (?)")
	handleException(err)

	res, err := query.Exec(sneaker.Name)
	id, err := res.LastInsertId()
	handleException(err)

	sneaker.ID = strconv.FormatInt(id, 10)

	json.NewEncoder(w).Encode(sneaker)
}

func DeleteSneaker(w http.ResponseWriter, r*http.Request) {
	params := mux.Vars(r)

	query, err := db.Prepare("DELETE FROM sneakers WHERE id = ?")
	handleException(err)

	res, err := query.Exec(params["id"])
	handleException(err)

	_ = res
}

func CreateTrueToSize(w http.ResponseWriter, r*http.Request) {
	params := mux.Vars(r)
	size := TrueToSize{SneakerID: params["id"] }
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &size)

	query, err := db.Prepare("INSERT INTO true_to_size (sneaker_id, size) VALUES (?,?)")
	handleException(err)

	res, err := query.Exec(size.SneakerID, size.Size)
	id, err := res.LastInsertId()
	handleException(err)

	size.ID = strconv.FormatInt(id, 10)

	json.NewEncoder(w).Encode(size)
}

func handleException(exception error) {
	if exception != nil {
		panic(exception)
	}
}