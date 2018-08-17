package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

/*
Test with this curl command:
curl -H "Content-Type: application/json" -d {\"Sum\":100} http://localhost:8080/transactions
*/
func TransactionCheck(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(jsonErr{Code: http.StatusBadRequest, Text: err.Error()})
		return
	}
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(jsonErr{Code: http.StatusBadRequest, Text: err.Error()})
		return
	}
	bodyString := string(body)
	if err := json.Unmarshal([]byte(bodyString), &transaction); err != nil {
		w.WriteHeader(422) // unprocessable entity
		json.NewEncoder(w).Encode(jsonErr{Code: 422, Text: err.Error()})
		return
	}
	if !СompletenessСheckTransaction(bodyString) {
		w.WriteHeader(422) // unprocessable entity
		json.NewEncoder(w).Encode(jsonErr{Code: 422, Text: "Invalid Data"})
		return
	}
	t := PredictTransaction(transaction)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
