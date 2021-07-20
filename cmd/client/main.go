package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/matheusmosca/beeserve"
)

const (
	port = 3000
	host = "localhost"
)

type Account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Document  string    `json:"document"`
	CreatedAT time.Time `json:"created_at"`
}

var accounts = []Account{
	{
		ID:        uuid.NewString(),
		Name:      "Maria",
		Document:  "dfkc02323",
		CreatedAT: time.Now(),
	},
	{
		ID:        uuid.NewString(),
		Name:      "José",
		Document:  "dfair21312",
		CreatedAT: time.Now(),
	},
}

func main() {
	address := fmt.Sprintf("%s:%d", host, port)
	fmt.Println("oi")

	metricsClient := beeserve.NewClient("accounts-app", "localhost", 8000)

	r := mux.NewRouter()
	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.Handle("/accounts", beeserve.WithMetrics(http.HandlerFunc(createAccount), metricsClient)).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(address, v1))
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	var account Account
	w.WriteHeader(http.StatusCreated)
	time.Sleep(time.Second * 1)
	json.NewDecoder(r.Body).Decode(&account)
}
