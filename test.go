package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// СompletenessСheckTransaction
//
//Format Json
func TestCheckTransInvalidJson(t *testing.T) {
	total := СompletenessСheckTransaction(`{"Sum": 5`)
	assert.NotNil(t, total, "The `total` should not be `nil`")
	assert.Equal(t, false, total, "Expecting `false`")
}

//KeysJson
func TestCheckTransNotEnoughKeysJson(t *testing.T) {
	total := СompletenessСheckTransaction(`{"Sum": 5}`)
	assert.NotNil(t, total, "The `total` should not be `nil`")
	assert.Equal(t, false, total, "Expecting `false`")
}
func TestCheckTransNotValidKeysJson(t *testing.T) {
	total := СompletenessСheckTransaction(`{"S": 5, "YearEnd": 2021, "Hour": 4.5, "WeekDay": 4, "HaveDeviceId": 0, "EuropeAsiaCountryIp": 0}`)
	assert.NotNil(t, total, "The `total` should not be `nil`")
	assert.Equal(t, false, total, "Expecting `false`")
}
func TestCheckTransValidKeysJson(t *testing.T) {
	total := СompletenessСheckTransaction(`{"Sum": 5, "YearEnd": 2021, "Hour": 4.5, "WeekDay": 4, "HaveDeviceId": 0, "EuropeAsiaCountryIp": 0}`)
	assert.NotNil(t, total, "The `total` should not be `nil`")
	assert.Equal(t, false, total, "Expecting `false`")
}

//Todo

//
// Requests
//
/*
func CreateEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Item Created"))
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", CreateEndpoint).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Router() *mux.Router {
	router := NewRouter()
	router.HandleFunc("/transactions", CreateEndpoint).Methods("POST")
	return router
}


func TestCreate(t *testing.T) {
	trans := &Transaction{
		Sum:                 0.01,
		YearEnd:             2021,
		Hour:                4.5,
		WeekDay:             4,
		HaveDeviceId:        0,
		EuropeAsiaCountryIp: 0,
	}
	jsonTrans, _ := json.Marshal(trans)
	request, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(jsonTrans))
	response := httptest.NewRecorder()
	NewRouter().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
*/
